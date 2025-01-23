package datacenter

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gofly/app/common/ffmpegtool"
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"net/http"
	"os"
	"strings"
	"time"
)

// 文件管理
type Upfile struct{ NoNeedAuths []string }

func init() {
	fpath := Upfile{NoNeedAuths: []string{"*"}}
	gf.Register(&fpath, fpath)
}

// b端上传文件总接口
func (api *Upfile) Upload(c *gf.GinCtx) {
	filetype := c.DefaultPostForm("filetype", "image") //文件类型
	uppath := "/business/datacenter/upfile/saveFile"
	if filetype == "image" {
		uppath = gf.GetUploadURL(uppath, "uploadpath_read")
	}
	//重定向到内部
	c.Redirect(http.StatusPermanentRedirect, uppath)
}

// 保存本地附件
func (api *Upfile) SaveFile(c *gf.GinCtx) {
	businessID := c.GetInt64("businessID") //当前商户ID
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
	defer fileContent.Close()
	var byteContainer []byte
	byteContainer = make([]byte, 1000000)
	fileContent.Read(byteContainer)
	m_d5 := md5.New()
	m_d5.Write(byteContainer)
	sha1_str := hex.EncodeToString(m_d5.Sum(nil))
	//查找该用户是否传过
	attachment, _ := gf.Model("business_attachment").Where("business_id", businessID).
		Where("sha1", sha1_str).Fields("id,pid,name,title,type,url,filesize,mimetype,cover_url as cover").Find()
	if attachment != nil { //文件是否已经存在
		//更新到最前面
		maxId, _ := gf.Model("business_attachment").Where("business_id", businessID).Order("weigh desc").Value("id")
		if maxId != nil {
			gf.Model("business_attachment").Data(map[string]interface{}{"weigh": maxId.Int() + 1, "pid": pid}).Where("id", attachment["id"]).Update()
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
			} else if filetype == "image" { //图片
				ftype = 0
			} else {
				ftype = 5
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
			getdata, _ := gf.Model("business_attachment").Where("id", file_id).Fields("id,pid,name,title,type,url,filesize,mimetype,cover_url as cover").Find()
			gf.Success().SetMsg("上传成功").SetData(getdata).Regin(c)
		}
	}
}
