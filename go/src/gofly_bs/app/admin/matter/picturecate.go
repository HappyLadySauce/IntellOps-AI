package matter

import (
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gmap"
)

type Picturecate struct{ NoNeedAuths []string }

func init() {
	fpath := Picturecate{NoNeedAuths: []string{"getCate", "GetList"}}
	gf.Register(&fpath, fpath)
}

// 获取列表
func (api *Picturecate) GetList(c *gf.GinCtx) {
	pageNo := gconv.Int(c.DefaultQuery("page", "1"))
	pageSize := gconv.Int(c.DefaultQuery("pageSize", "10"))
	//搜索添条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	if name, ok := param["name"]; ok && name != "" {
		whereMap.Set("name like ?", "%"+gconv.String(name)+"%")
	}
	MDB := gf.Model("common_picture_cate").Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Page(pageNo, pageSize).Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		gf.Success().SetMsg("获取全部列表").SetData(gf.Map{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}).Regin(c)
	}
}

// 获取分类列表
func (api *Picturecate) GetCate(c *gf.GinCtx) {
	list, err := gf.Model("common_picture_cate").Where("status", 0).Fields("id,name,type").Order("weigh desc,id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		gf.Success().SetMsg("获取选择列表").SetData(list).Regin(c)
	}
}

// 保存
func (api *Picturecate) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	var f_id = gf.GetEditId(param["id"])
	if f_id == 0 {
		delete(param, "id")
		addId, err := gf.Model("common_picture_cate").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
		} else {
			if addId != 0 {
				gf.Model("common_picture_cate").Data(map[string]interface{}{"weigh": addId}).Where("id", addId).Update()
			}
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("common_picture_cate").Data(param).Where("id", f_id).Update()
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 更新状态
func (api *Picturecate) UpStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("common_picture_cate").Where("id", param["id"]).Data(map[string]interface{}{"status": param["status"]}).Update()
	if err != nil {
		gf.Failed().SetMsg("更新失败！").SetData(err).Regin(c)
	} else {
		msg := "更新成功！"
		if res2 == nil {
			msg = "暂无数据更新"
		}
		gf.Success().SetMsg(msg).SetData(res2).Regin(c)
	}
}

// 删除
func (api *Picturecate) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("common_picture_cate").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
	}
}
