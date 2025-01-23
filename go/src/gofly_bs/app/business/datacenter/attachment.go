package datacenter

import (
	"fmt"
	"gofly/utils/gf"
	"gofly/utils/tools/gmap"
	"gofly/utils/tools/gvar"
	"net/http"
	"strings"
)

type Attachment struct{ NoNeedAuths []string }

func init() {
	fpath := Attachment{NoNeedAuths: []string{"*"}}
	gf.Register(&fpath, fpath)
}

// 获取分类列表
func (api *Attachment) GetPictureCate(c *gf.GinCtx) {
	list, err := gf.Model("common_picture_cate").Where("status", 0).Fields("id,name,type").Order("weigh desc,id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		gf.Success().SetMsg("获取选择列表").SetData(list).Regin(c)
	}
}

// 获取图片库列表-图库以本地为主
func (api *Attachment) GetPicture(c *gf.GinCtx) {
	pageNo := gf.Int(c.DefaultQuery("page", "1"))
	pageSize := gf.Int(c.DefaultQuery("pageSize", "10"))
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	whereMap.Set("status", 0)
	if cid, ok := param["cid"]; ok && gf.Int(cid) != 0 {
		whereMap.Set("cid", cid)
	}
	if types, ok := param["type"]; ok {
		whereMap.Set("type", types)
	}
	if title, ok := param["title"]; ok && title != "" {
		whereMap.Set("title like ?", "%"+gf.String(title)+"%")
	}
	MDB := gf.Model("common_picture").Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Fields("id,cid,url,type,title,mimetype,cover_url,createtime").Page(pageNo, pageSize).Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		gf.Success().SetMsg("获取全部图库列表").SetData(gf.Map{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}).Regin(c)
	}
}

// 保存
func (api *Attachment) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	businessID := c.GetInt64("businessID") //当前商户ID
	var f_id = gf.GetEditId(param["id"])
	if f_id == 0 {
		param["business_id"] = businessID
		getcount, _ := gf.Model("business_attachment").Where("business_id", businessID).Where("pid", param["pid"]).Where("title like ?", fmt.Sprintf("%s%v%s", "%", param["title"], "%")).Count()
		param["title"] = fmt.Sprintf("%s%v", param["title"], getcount+1)
		addId, err := gf.Model("business_attachment").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").Regin(c)
		} else {
			//更新排序
			gf.Model("business_attachment").Data(map[string]interface{}{"weigh": addId}).Where("id", addId).Update()
			getdata, _ := gf.Model("business_attachment").Where("id", addId).Fields("id,pid,name,title,type,url,filesize,mimetype,storage").Select()
			gf.Success().SetMsg("添加成功！").SetData(getdata).Regin(c)
		}
	} else {
		res, err := gf.Model("business_attachment").Data(param).Where("id", f_id).Update()
		if err != nil {
			gf.Failed().SetMsg("更新失败").Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 删除
func (api *Attachment) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	thrid, _ := gf.Model("business_attachment").WhereIn("id", param["ids"]).Where("storage !=?", "local").Count()
	file_list, _ := gf.Model("business_attachment").WhereIn("id", param["ids"]).Where("storage", "local").Where("type != ?", 1).Array("url")
	res2, err := gf.Model("business_attachment").WhereIn("id", param["ids"]).Where("storage", "local").Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").Regin(c)
	} else {
		if file_list != nil {
			gf.Del_file(file_list)
		}
		if thrid > 0 {
			uppath := gf.GetUploadURL("", "delpath_read")
			if uppath != "" {
				//重定向到删除云存储附件
				c.Redirect(http.StatusPermanentRedirect, uppath)
				return
			}
		}
		gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
	}
}

// 删除文件夹
func (api *Attachment) DelDir(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("business_attachment").Where("id", param["id"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").Regin(c)
	} else {
		delCycle(param["id"])
		gf.Success().SetMsg("文件夹成功！").SetData(res2).Regin(c)
	}
}

// 循环删除子文件
func delCycle(id interface{}) {
	filedata, _ := gf.Model("business_attachment").Where("pid", id).Fields("id,type").Select()
	if !filedata.IsEmpty() && len(filedata) > 0 {
		for _, val := range filedata {
			delCycle(val["id"])
			if val["type"].Int64() == 0 {
				file_path, _ := gf.Model("business_attachment").Where("id", val["id"]).Value("url")
				if file_path != nil {
					gf.DelOneFile(file_path.String())
				}
			}
			gf.Model("business_attachment").Where("id", val["id"]).Delete()
		}
	}
}

// 获取我的附件
func (api *Attachment) GetMyFiles(c *gf.GinCtx) {
	searchword := c.DefaultQuery("searchword", "")
	filetype := c.DefaultQuery("filetype", "image")
	pid := c.DefaultQuery("pid", "0")
	businessID := c.GetInt64("businessID") //当前商户ID
	whereMap := gf.Model("business_attachment").Where("business_id", businessID).Where("pid", pid)
	if searchword != "" {
		whereMap.Where("title like ?", "%"+searchword+"%")
	}
	if filetype == "video" {
		whereMap.WhereIn("type", []interface{}{1, 2})
	} else if filetype == "audio" {
		whereMap.WhereIn("type", []interface{}{1, 3})
	} else if filetype == "file" {
		whereMap.WhereIn("type", []interface{}{1, 4})
	} else { //默认图片
		whereMap.WhereIn("type", []interface{}{0, 1})
	}
	list, err := whereMap.Fields("id,pid,name,title,type,url,filesize,mimetype,storage,cover_url,is_common").Order("type desc,weigh desc,id desc").Select()
	if err != nil {
		gf.Failed().SetMsg("加载数据失败").Regin(c)
	} else {
		rooturl := gf.GetMainURLLocal()
		for _, val := range list {
			if _, ok := val["cover_url"]; ok && val["cover_url"].String() != "" && !strings.Contains(val["cover_url"].String(), "http") && rooturl != "" {
				val["cover_url"] = gvar.New(rooturl + val["cover_url"].String())
			}
		}
		common_lisr, _ := gf.Model("business_attachment").Where("pid", pid).Where("is_common", 1).Fields("id,pid,name,title,type,url,filesize,mimetype,storage,cover_url,is_common").Order("type desc,weigh desc,id desc").Select()
		if list != nil {
			list = append(common_lisr, list...)
		} else {
			list = common_lisr
		}
		var totalCount int64
		//获取目录菜单
		allids := getAllParentIds(pid)
		allids = append(allids, pid)
		dirmenu, _ := gf.Model("business_attachment").WhereIn("id", allids).Fields("id,pid,title").Select()
		hasegallery := false
		dielddata, err := gf.DB().GetCore().HasTable("common_picture")
		if err == nil && dielddata {
			hasegallery = true
		}
		resdata := gf.Map{
			"total":       totalCount,
			"dirmenu":     dirmenu,
			"allids":      allids,
			"items":       list,
			"hasegallery": hasegallery, //是否存在图库
		}
		gf.Success().SetMsg("获取附件列表").SetData(resdata).Regin(c)
	}
}

// 工具
func getAllParentIds(id interface{}) []interface{} {
	var parent_ids []interface{}
	parent_id, _ := gf.Model("business_attachment").Where("id", id).Value("pid")
	if parent_id != nil {
		parent_ids = append(parent_ids, parent_id)
		parent_ids = append(parent_ids, getAllParentIds(parent_id)...)
	}
	return parent_ids
}

// 更新图片目录
func (api *Attachment) UpImgPid(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res, err := gf.Model("business_attachment").Where("id", param["imgid"]).Data(map[string]interface{}{"pid": param["pid"]}).Update()
	if err != nil {
		gf.Failed().SetMsg("更新失败！").SetData(err).Regin(c)
	} else {
		msg := "更新目录成功！"
		if res == nil {
			msg = "暂无目录更新"
		}
		gf.Success().SetMsg(msg).Regin(c)
	}
}
