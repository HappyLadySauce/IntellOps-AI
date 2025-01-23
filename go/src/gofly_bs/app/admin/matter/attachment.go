package matter

import (
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gmap"
	"gofly/utils/tools/gvar"
	"strings"
)

type Attachment struct{}

func init() {
	fpath := Attachment{}
	gf.Register(&fpath, fpath)
}

// 获取列表
func (api *Attachment) GetList(c *gf.GinCtx) {
	user, exists := gf.GetUserInfo(c) //当前用户
	if !exists {
		gf.Failed().SetMsg("登录失效").Regin(c)
		return
	}
	pageNo := gconv.Int(c.DefaultQuery("page", "1"))
	pageSize := gconv.Int(c.DefaultQuery("pageSize", "10"))
	//搜索添条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	whereMap.Set("business_id", user.BusinessID)
	if title, ok := param["title"]; ok && title != "" {
		whereMap.Set("title like ?", "%"+gconv.String(title)+"%")
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	MDB := gf.Model("attachment").Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Fields("id,url,imagewidth,imageheight,imagetype,filesize,mimetype,createtime,sha1,title,name,cover_url").Page(pageNo, pageSize).Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).SetData(err).Regin(c)
	} else {
		rooturl := gf.GetMainURLLocal()
		for _, val := range list {
			if _, ok := val["url"]; ok && val["url"].String() != "" && !strings.Contains(val["url"].String(), "http") {
				val["url"] = gvar.New(rooturl + val["url"].String())
			}
		}
		gf.Success().SetMsg("获取全部列表").SetData(gf.Map{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}).Regin(c)
	}
}

// 删除
func (api *Attachment) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	file_list, _ := gf.Model("attachment").WhereIn("id", param["ids"]).Array("url")
	res2, err := gf.Model("attachment").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		if file_list != nil {
			gf.Del_file(file_list)
		}
		gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
	}
}
