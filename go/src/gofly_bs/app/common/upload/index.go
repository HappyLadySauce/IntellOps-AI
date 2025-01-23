package upload

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gofly/app/common/ffmpegtool"
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gvar"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 通用文件上传-小程序、app、h5等
type Index struct{ NoNeedAuths []string }

func init() {
	fpath := Index{NoNeedAuths: []string{"*"}}
	gf.Register(&fpath, fpath)
}

// 业务端通用上传文件总接口
// 请求头添加 Businessid=当sass系统时记录(默认1账号)，filetype=附件类型(默认图片)
func (api *Index) Action(c *gf.GinCtx) {
	filetype := c.DefaultPostForm("filetype", "image") //文件类型默认图片
	uppath := "/common/upload/upFile"
	if filetype == "image" || filetype == "doc" { //图片和文件类可以上传三方云存储
		uppath = gf.GetUploadURL(uppath, "uploadpath_read")
	}
	//重定向到内部
	c.Redirect(http.StatusPermanentRedirect, uppath)
}

// 上传业务附件
func (api *Index) UpFile(c *gf.GinCtx) {
	var businessID any = c.GetHeader("Businessid") //从请求头获取businessID判断是那个服务端传过来
	if businessID == "" {                          //找不到在去登录token获取
		businessID, _ = c.Get("businessID") //当前用户businessID(saas账号ID)
	}
	if businessID == "" { //找不到在去登录token获取
		businessID = 1 //默认单服务系统
	}
	//判断存储空间是否超出
	var usesize interface{}
	var fileSize interface{}
	usesize, _ = gf.Model("business_attachment").Where("business_id", businessID).Where("type", 0).Sum("filesize")
	if usesize == nil {
		usesize = 0
	}
	fileSize, _ = gf.Model("business_account").Where("id", businessID).Value("fileSize")
	if fileSize == nil {
		fileSize = 0
	}
	if gconv.Int(usesize) >= gconv.Int(fileSize) {
		gf.Failed().SetMsg("您的存储空间已满，请您先去购买存储空间！").Regin(c)
		return
	}
	// 单个文件
	pid := c.DefaultPostForm("pid", "1")
	filetype := c.DefaultPostForm("filetype", "image") //文件类型
	file, err := c.FormFile("file")
	if err != nil {
		gf.Failed().SetMsg("获取数据失败，").SetData(err).Regin(c)
		return
	}
	nowTime := time.Now().Unix() //当前时间
	rooturl := gf.GetMainURL()
	//判断文件是否已经传过
	fileContent, _ := file.Open()
	defer fileContent.Close()
	var byteContainer []byte
	byteContainer = make([]byte, 1000000)
	fileContent.Read(byteContainer)
	m_d5 := md5.New()
	m_d5.Write(byteContainer)
	sha1_str := hex.EncodeToString(m_d5.Sum(nil))
	//查找该用户是否传过
	attachment, _ := gf.Model("business_attachment").Where("business_id", businessID).
		Where("sha1", sha1_str).Fields("id,pid,name,title,type,url,filesize,mimetype,cover_url").Find()
	if attachment != nil { //文件是否已经存在
		//更新到最前面
		maxId, _ := gf.Model("business_attachment").Where("business_id", businessID).Order("weigh desc").Value("id")
		if maxId != nil {
			gf.Model("business_attachment").Data(map[string]interface{}{"weigh": maxId.Int() + 1, "pid": pid}).Where("id", attachment["id"]).Update()
		}
		attachment["rooturl"] = gvar.New(rooturl)
		attachment["fullpath"] = gvar.New(rooturl + attachment["url"].String())
		gf.Success().SetMsg("文件已上传").SetData(attachment).Regin(c)
	} else {
		file_path := fmt.Sprintf("%s%s%s", "resource/uploads/", time.Now().Format("20060102"), "/")
		//如果没有filepath文件目录就创建一个
		if _, err := os.Stat(file_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(file_path, os.ModePerm)
			}
		}
		//上传到的路径
		filename_arr := strings.Split(file.Filename, ".")
		//重新名片-lunix系统不支持中文
		name_str := gf.Md5Str(fmt.Sprintf("%v%s", nowTime, filename_arr[0]))   //组装文件保存名字
		file_Filename := fmt.Sprintf("%s%s%s", name_str, ".", filename_arr[1]) //文件加.后缀
		path := file_path + file_Filename
		// 上传文件到指定的目录
		err = c.SaveUploadedFile(file, path)
		if err != nil { //上传失败
			gf.Failed().SetMsg("上传失败").SetData(err).Regin(c)
		} else { //上传成功
			//保存数据
			var ftype int64 = 0
			var cover_url string = ""
			if filetype == "video" {
				ftype = 2
				videopath := fmt.Sprintf("./%s", path)
				pathroot := strings.Split(path, ".")
				imgpath := fmt.Sprintf("./%s", pathroot[0])
				fname, err := ffmpegtool.GetSnapshot(videopath, imgpath, 1)
				if err == nil {
					cover_url = fname
				}
			} else if filetype == "audio" { //音频
				ftype = 3
			} else if filetype == "file" { //附件类
				ftype = 4
			}
			Insertdata := map[string]interface{}{
				"business_id": businessID,
				"type":        ftype,
				"pid":         pid,
				"sha1":        sha1_str,
				"title":       filename_arr[0],
				"name":        file.Filename,
				"url":         path,
				"cover_url":   cover_url, //视频封面
				"storage":     "local",
				"createtime":  nowTime,
				"filesize":    file.Size,
				"mimetype":    file.Header["Content-Type"][0],
			}
			//保存数据
			file_id, _ := gf.Model("business_attachment").Data(Insertdata).InsertAndGetId()
			//更新排序
			gf.Model("business_attachment").Data(map[string]interface{}{"weigh": file_id}).Where("id", file_id).Update()
			//返回数据
			getdata, _ := gf.Model("business_attachment").Where("id", file_id).Fields("id,pid,name,title,type,url,filesize,mimetype,cover_url").Find()
			getdata["rooturl"] = gvar.New(rooturl)
			getdata["fullpath"] = gvar.New(rooturl + getdata["url"].String())
			gf.Success().SetMsg("上传成功").SetData(getdata).Regin(c)
		}
	}
}

// 3.上传单文件不验证
func (api *Index) FileNov(c *gf.GinCtx) {
	// 单个文件
	file, err := c.FormFile("file")
	if err != nil {
		gf.Failed().SetMsg("获取数据失败").SetData(err).Regin(c)
		return
	}
	nowTime := time.Now().Unix() //当前时间
	//时间查询-获取当天时间
	day_time := time.Now().Format("2006-01-02")
	//文件唯一性
	file_uniname := fmt.Sprintf("%s%s%v", file.Filename, day_time, "")
	sha1_str := gf.Md5Str(file_uniname)
	rooturl := gf.GetMainURL()
	file_path := fmt.Sprintf("%s%s%s", "resource/uploads/", time.Now().Format("20060102"), "/")
	//如果没有filepath文件目录就创建一个
	if _, err := os.Stat(file_path); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(file_path, os.ModePerm)
		}
	}
	//上传到的路径
	filename_arr := strings.Split(file.Filename, ".")
	name_str := gf.Md5Str(fmt.Sprintf("%v%s", nowTime, filename_arr[0]))   //组装文件保存名字
	file_Filename := fmt.Sprintf("%s%s%s", name_str, ".", filename_arr[1]) //文件加.后缀
	path := file_path + file_Filename
	// 上传文件到指定的目录
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		gf.Failed().SetMsg("上传失败").SetData(err).Regin(c)
	} else {
		//保存数据
		Insertdata := map[string]interface{}{
			"business_id": 1,
			"type":        4,
			"pid":         1,
			"sha1":        sha1_str,
			"title":       filename_arr[0],
			"name":        file.Filename,
			"url":         path,
			"storage":     "local",
			"createtime":  nowTime,
			"filesize":    file.Size,
			"mimetype":    file.Header["Content-Type"][0],
		}
		//保存数据
		gf.Model("business_attachment").Data(Insertdata).InsertAndGetId()
		c.JSON(200, gin.H{
			"name":     file.Filename,
			"status":   "done",
			"url":      fmt.Sprintf("%s%s", rooturl, path),
			"response": "上传成功",
			"time":     nowTime,
		})
	}
}

// 编辑器保存第三方图片到本地
func (api *Index) ThirdImage(c *gf.GinCtx) {
	businessID, _ := c.Get("businessID") //当前用户businessID(saas账号ID)
	if businessID == "" {                //找不到在去登录token获取
		businessID = 1 //默认单服务系统
	}
	params, _ := gf.RequestParam(c)
	if url, ok := params["url"]; !ok || url == "" {
		c.JSON(200, gin.H{
			"code":   400,
			"result": false,
			"data": map[string]interface{}{
				"url": "",
			},
			"message": "地址无效",
		})
	} else {
		file_path := fmt.Sprintf("%s%s%s", "resource/uploads/", time.Now().Format("20060102"), "/")
		if _, err := os.Stat(file_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(file_path, os.ModePerm)
			}
		}
		nowTime := time.Now().Unix() //当前时间
		localPicName := fmt.Sprintf("%vthir_%v", file_path, nowTime)
		imgtype, err := gf.DownPic(gconv.String(params["url"]), localPicName)
		imgurl := fmt.Sprintf("%s%s", localPicName, imgtype)
		Insertdata := map[string]interface{}{
			"business_id": businessID,
			"type":        0,
			"pid":         0,
			"title":       fmt.Sprintf("thir_%v%v", nowTime, imgtype),
			"name":        fmt.Sprintf("thir_%v", nowTime),
			"url":         imgurl,
			"storage":     "local",
			"createtime":  nowTime,
			"mimetype":    "image/png",
		}
		//保存数据
		gf.Model("business_attachment").Data(Insertdata).InsertAndGetId()
		c.JSON(200, gin.H{
			"code":    200,
			"result":  true,
			"err":     err,
			"status":  "done",
			"url":     gf.GetFullPath(imgurl),
			"message": "上传成功",
		})
	}
}

// 上传admin端图片 系统默认附件attachment表，不限制空间
func (api *Index) Fileadmin(c *gf.GinCtx) {
	// 单个文件
	pid := c.DefaultPostForm("pid", "")
	filetype := c.DefaultPostForm("filetype", "image") //文件类型
	file, err := c.FormFile("file")
	if err != nil {
		gf.Failed().SetMsg("获取数据失败，").SetData(err).Regin(c)
		return
	}
	nowTime := time.Now().Unix() //当前时间
	//判断文件是否已经传过
	fileContent, _ := file.Open()
	var byteContainer []byte
	byteContainer = make([]byte, 1000000)
	fileContent.Read(byteContainer)
	m_d5 := md5.New()
	m_d5.Write(byteContainer)
	sha1_str := hex.EncodeToString(m_d5.Sum(nil))
	//查找该用户是否传过
	attachment, _ := gf.Model("attachment").
		Where("sha1", sha1_str).Fields("id,pid,name,title,type,url,filesize,mimetype,cover_url").Find()
	if attachment != nil { //文件是否已经存在
		//更新到最前面
		maxId, _ := gf.Model("attachment").Order("weigh desc").Value("id")
		if maxId != nil {
			gf.Model("attachment").Data(map[string]interface{}{"weigh": maxId.Int() + 1, "pid": pid}).Where("id", attachment["id"]).Update()
		}
		gf.Success().SetMsg("文件已上传").SetData(attachment).Regin(c)
	} else {
		file_path := fmt.Sprintf("%s%s%s", "resource/uploads/", time.Now().Format("20060102"), "/")
		//如果没有filepath文件目录就创建一个
		if _, err := os.Stat(file_path); err != nil {
			if !os.IsExist(err) {
				os.MkdirAll(file_path, os.ModePerm)
			}
		}
		//上传到的路径
		filename_arr := strings.Split(file.Filename, ".")
		//重新名片-lunix系统不支持中文
		name_str := gf.Md5Str(fmt.Sprintf("%v%s", nowTime, filename_arr[0]))   //组装文件保存名字
		file_Filename := fmt.Sprintf("%s%s%s", name_str, ".", filename_arr[1]) //文件加.后缀
		path := file_path + file_Filename
		// 上传文件到指定的目录
		err = c.SaveUploadedFile(file, path)
		if err != nil { //上传失败
			gf.Failed().SetMsg("上传失败").SetData(err).Regin(c)
		} else { //上传成功
			//保存数据
			var ftype int64 = 0
			var cover_url string = ""
			if filetype == "video" {
				ftype = 2
				videopath := fmt.Sprintf("./%s", path)
				pathroot := strings.Split(path, ".")
				imgpath := fmt.Sprintf("./%s", pathroot[0])
				fname, err := ffmpegtool.GetSnapshot(videopath, imgpath, 1)
				if err == nil {
					cover_url = fname
				}
			} else if filetype == "audio" { //音频
				ftype = 3
			} else if filetype == "file" { //附件类
				ftype = 4
			}
			Insertdata := map[string]interface{}{
				"type":      ftype,
				"pid":       pid,
				"sha1":      sha1_str,
				"title":     filename_arr[0],
				"name":      file.Filename,
				"url":       path,
				"cover_url": cover_url, //视频封面
				"filesize":  file.Size,
				"mimetype":  file.Header["Content-Type"][0],
			}
			//保存数据
			file_id, _ := gf.Model("attachment").Data(Insertdata).InsertAndGetId()
			//更新排序
			gf.Model("attachment").Data(map[string]interface{}{"weigh": file_id}).Where("id", file_id).Update()
			//返回数据
			getdata, _ := gf.Model("attachment").Where("id", file_id).Fields("id,pid,name,title,type,url,filesize,mimetype").Find()
			gf.Success().SetMsg("上传成功").SetData(getdata).Regin(c)
		}
	}
}
