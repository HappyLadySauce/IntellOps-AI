package matter

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gmap"
	"gofly/utils/tools/gvar"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Picture struct{ NoNeedAuths []string }

func init() {
	fpath := Picture{NoNeedAuths: []string{"uploadFile"}}
	gf.Register(&fpath, fpath)
}

// 获取列表
func (api *Picture) GetList(c *gf.GinCtx) {
	pageNo := gconv.Int(c.DefaultQuery("page", "1"))
	pageSize := gconv.Int(c.DefaultQuery("pageSize", "10"))
	//搜索添条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	if title, ok := param["title"]; ok && title != "" {
		whereMap.Set("title like ?", "%"+gconv.String(title)+"%")
	}
	if status, ok := param["status"]; ok && status != "" {
		whereMap.Set("status", status)
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	MDB := gf.Model("common_picture").Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Page(pageNo, pageSize).Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		for _, val := range list {
			if _, ok := val["cid"]; ok {
				catename, _ := gf.Model("common_picture_cate").Where("id", val["cid"]).Value("name")
				val["catename"] = catename
			}
		}
		gf.Success().SetMsg("获取全部列表").SetData(gf.Map{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}).Regin(c)
	}
}

// 保存
func (api *Picture) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	delete(param, "catename")
	res, err := gf.Model("common_picture").Data(param).Where("id", param["id"]).Update()
	if err != nil {
		gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
	}
}

// 更新状态
func (api *Picture) UpStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res, err := gf.Model("common_picture").Where("id", param["id"]).Data(map[string]interface{}{"status": param["status"]}).Update()
	if err != nil {
		gf.Failed().SetMsg("更新失败！").SetData(err).Regin(c)
	} else {
		msg := "更新成功！"
		if res == nil {
			msg = "暂无数据更新"
		}
		gf.Success().SetMsg(msg).Regin(c)
	}
}

// 删除
func (api *Picture) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	file_list, _ := gf.Model("common_picture").WhereIn("id", param["ids"]).Array("url")
	res2, err := gf.Model("common_picture").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		if file_list != nil {
			gf.Del_file(file_list)
		}
		gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
	}
}

// 上传图片
func (api *Picture) UploadFile(c *gf.GinCtx) {
	// 单个文件
	cid := c.DefaultPostForm("cid", "")
	typeId := c.DefaultPostForm("type", "")
	Id := c.DefaultPostForm("id", "0")
	file, err := c.FormFile("file")
	if err != nil {
		gf.Failed().SetMsg("获取数据失败，").SetData(err).Regin(c)
		return
	}
	nowTime := time.Now().Unix()      //当前时间
	user, exists := gf.GetUserInfo(c) //当前用户
	if !exists {
		gf.Failed().SetMsg("登录失效").Regin(c)
		return
	}
	//判断文件是否已经传过
	fileContent, _ := file.Open()
	var byteContainer []byte
	byteContainer = make([]byte, 1000000)
	fileContent.Read(byteContainer)
	m_d5 := md5.New()
	m_d5.Write(byteContainer)
	sha1_str := hex.EncodeToString(m_d5.Sum(nil))
	//查找该用户是否传过
	attachment, _ := gf.Model("common_picture").Where("uid", user.ID).
		Where("sha1", sha1_str).Fields("id,name,title,url,filesize,mimetype,storage").Find()
	if attachment != nil { //文件是否已经存在
		//更新到最前面
		var nid interface{}
		if Id != "0" {
			gf.Model("common_picture").Data(map[string]interface{}{"title": attachment["title"], "name": attachment["name"], "url": attachment["url"]}).Where("id", Id).Update()
			nid = Id
		} else {
			delete(attachment, "id")
			attachment["cid"] = gvar.New(cid)
			attachment["type"] = gvar.New(typeId)
			file_id, _ := gf.Model("common_picture").Data(attachment).InsertAndGetId()
			nid = file_id
			//更新排序
			gf.Model("common_picture").Data(map[string]interface{}{"weigh": file_id}).Where("id", file_id).Update()
		}
		gf.Success().SetMsg("文件已上传").SetData(gf.Map{"id": nid, "title": attachment["title"], "url": attachment["url"]}).Regin(c)
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
		name_str := gf.Md5(fmt.Sprintf("%v%s", nowTime, filename_arr[0]))      //组装文件保存名字
		file_Filename := fmt.Sprintf("%s%s%s", name_str, ".", filename_arr[1]) //文件加.后缀
		path := file_path + file_Filename
		// 上传文件到指定的目录
		err = c.SaveUploadedFile(file, path)
		if err != nil { //上传失败
			c.JSON(200, gin.H{
				"uid":      sha1_str,
				"name":     file.Filename,
				"status":   "error",
				"response": "上传失败",
				"time":     nowTime,
			})
		} else { //上传成功
			//保存数据
			Insertdata := map[string]interface{}{
				"uid":        user.ID,
				"type":       typeId,
				"cid":        cid,
				"sha1":       sha1_str,
				"title":      filename_arr[0],
				"name":       file.Filename,
				"url":        path,
				"createtime": nowTime,
				"filesize":   file.Size,
				"mimetype":   file.Header["Content-Type"][0],
			}
			//保存数据
			var nid interface{}
			if Id != "0" {
				gf.Model("common_picture").Data(Insertdata).Where("id", Id).Update()
				nid = Id
			} else {
				file_id, _ := gf.Model("common_picture").Data(Insertdata).InsertAndGetId()
				nid = file_id
				//更新排序
				gf.Model("common_picture").Data(map[string]interface{}{"weigh": file_id}).Where("id", file_id).Update()
			}
			//返回数据
			gf.Success().SetMsg("上传成功!").SetData(gf.Map{"id": nid, "title": Insertdata["title"], "url": Insertdata["url"].(string)}).Regin(c)
		}
	}
}
