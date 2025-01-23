package datacenter

import (
	"fmt"
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gvar"
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

// 获取图片库列表
func (api *Attachment) GetPicture(c *gf.GinCtx) {
	cid := c.DefaultQuery("cid", "0")
	types := c.DefaultQuery("type", "0")
	title := c.DefaultQuery("searchword", "")
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo := gconv.Int(page)
	pageSize := gconv.Int(_pageSize)
	MDB := gf.Model("common_picture").
		Fields("id,cid,url,type,title,mimetype,cover_url,createtime").Where("status", 0).Where("type", types)
	if cid != "0" {
		MDB.Where("cid", cid)
	}
	if title != "" {
		MDB.Where("title like ?", "%"+title+"%")
	}
	list, err := MDB.Page(pageNo, pageSize).Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		rooturl := gf.GetMainURL()
		for _, val := range list {
			if _, ok := val["image"]; ok && val["image"].String() != "" && !strings.Contains(val["image"].String(), "http") {
				val["image"] = gvar.New(rooturl + val["image"].String())
			}
		}
		var totalCount int64
		gf.Success().SetMsg("获取全部列表").SetData(gf.Map{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}).Regin(c)
	}
}

// 保存
func (api *Attachment) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	var f_id = gf.GetEditId(param["id"])
	if f_id == 0 {
		param["uid"], _ = c.Get("userID")
		getcount, _ := gf.Model("attachment").Where("pid", param["pid"]).Where("title like ?", fmt.Sprintf("%s%v%s", "%", param["title"], "%")).Count()
		param["title"] = fmt.Sprintf("%s%v", param["title"], getcount+1)
		addId, err := gf.Model("attachment").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").Regin(c)
		} else {
			//更新排序
			gf.Model("attachment").Data(map[string]interface{}{"weigh": addId}).Where("id", addId).Update()
			getdata, _ := gf.Model("attachment").Where("id", addId).Fields("id,pid,name,title,type,url,filesize,mimetype,storage").Select()
			gf.Success().SetMsg("添加成功！").SetData(getdata).Regin(c)
		}
	} else {
		res, err := gf.Model("attachment").
			Data(param).
			Where("id", f_id).
			Update()
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
	file_list, _ := gf.Model("attachment").WhereIn("id", param["ids"]).Where("type != ?", 1).Array("url")
	res2, err := gf.Model("attachment").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").Regin(c)
	} else {
		if file_list != nil {
			gf.Del_file(file_list)
		}
		gf.Success().SetMsg("删除成功！").SetData(res2).SetExdata(file_list).Regin(c)
	}
}

// 删除文件夹
func (api *Attachment) DelDir(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("attachment").Where("id", param["id"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").Regin(c)
	} else {
		delCycle(param["id"])
		gf.Success().SetMsg("文件夹成功！").SetData(res2).Regin(c)
	}
}

// 循环删除子文件
func delCycle(id interface{}) {
	filedata, _ := gf.Model("attachment").Where("pid", id).Fields("id,type").Select()
	if filedata != nil && len(filedata) > 0 {
		for _, val := range filedata {
			delCycle(val["id"])
			if val["type"].Int64() == 0 {
				file_path, _ := gf.Model("attachment").Where("id", val["id"]).Value("url")
				if file_path != nil {
					gf.DelOneFile(file_path.String())
				}
			}
			gf.Model("attachment").Where("id", val["id"]).Delete()
		}
	}
}

// 获取我的附件
func (api *Attachment) GetMyFiles(c *gf.GinCtx) {
	searchword := c.DefaultQuery("searchword", "")
	filetype := c.DefaultQuery("filetype", "image")
	pid := c.DefaultQuery("pid", "0")
	//当前用户
	user, exists := gf.GetUserInfo(c) //当前用户
	if !exists {
		gf.Failed().SetMsg("url地址为空").Regin(c)
		gf.Failed().SetMsg("登录失效").Regin(c)
		return
	}
	whereMap := gf.Model("attachment").Where("business_id", user.BusinessID).Where("pid", pid)
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
	list, err := whereMap.
		Fields("id,pid,name,title,type,url,filesize,mimetype,storage,cover_url,is_common").Order("type desc,weigh desc,id desc").Select()
	if err != nil {
		gf.Failed().SetMsg("加载数据失败").SetData(err).Regin(c)
	} else {
		for _, val := range list {
			if _, ok := val["cover_url"]; ok && val["cover_url"].String() != "" && !strings.Contains(val["cover_url"].String(), "http") {
				val["cover_url"] = gvar.New(gf.GetFullPath(val["cover_url"].String()))
			}
		}
		common_lisr, _ := gf.Model("attachment").Where("pid", pid).Where("is_common", 1).Fields("id,pid,name,title,type,url,filesize,mimetype,storage,cover_url,is_common").Order("type desc,weigh desc,id desc").Select()
		if list != nil {
			list = append(common_lisr, list...)
		} else {
			list = common_lisr
		}
		var totalCount int64
		//获取目录菜单
		allids := getAllParentIds(pid)
		allids = append(allids, pid)
		dirmenu, _ := gf.Model("attachment").WhereIn("id", allids).Fields("id,pid,title").Select()
		gf.Success().SetMsg("获取附件列表").SetData(gf.Map{
			"total":   totalCount,
			"dirmenu": dirmenu,
			"allids":  allids,
			"items":   list,
		}).Regin(c)
	}
}

// 工具
func getAllParentIds(id interface{}) []interface{} {
	var parent_ids []interface{}
	parent_id, _ := gf.Model("attachment").Where("id", id).Value("pid")
	if parent_id != nil {
		parent_ids = append(parent_ids, parent_id)
		parent_ids = append(parent_ids, getAllParentIds(parent_id)...)
	}
	return parent_ids
}

// 更新图片目录
func (api *Attachment) UpImgPid(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("attachment").Where("id", param["imgid"]).Data(map[string]interface{}{"pid": param["pid"]}).Update()
	if err != nil {
		gf.Failed().SetMsg("更新失败！").SetData(err).Regin(c)
	} else {
		msg := "更新目录成功！"
		if res2 == nil {
			msg = "暂无目录更新"
		}
		gf.Success().SetMsg(msg).SetData(res2).Regin(c)
	}
}
