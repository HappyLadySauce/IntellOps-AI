package createcode

import (
	"gofly/utils/gf"
	"gofly/utils/tools/gmap"
)

// 左侧分类菜单
type Product struct{}

func init() {
	fpath := Product{}
	gf.Register(&fpath, fpath)
}

// 获取列表
func (api *Product) GetList(c *gf.GinCtx) {
	pageNo := gf.Int(c.DefaultQuery("page", "1"))
	pageSize := gf.Int(c.DefaultQuery("pageSize", "10"))
	//搜索添条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()

	if cid, ok := param["cid"]; ok && gf.Int(cid) != 0 {
		cids := gf.CateAllChilId("createcode_product_cate", cid)
		whereMap.Set("cid In(?)", cids)
	}
	if title, ok := param["title"]; ok && title != "" {
		whereMap.Set("title", title)
	}
	if userType, ok := param["userType"]; ok && userType != "" {
		whereMap.Set("userType", userType)
	}
	if status, ok := param["status"]; ok && status != "" {
		whereMap.Set("status", status)
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	MDB := gf.Model("createcode_product").Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Fields("id,title,image,cid,num,price,sex,likeColor,userType,images,status,createtime,updatetime,des").Page(pageNo, pageSize).Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {

		for _, val := range list {
			val["cidName"] = gf.GetTalbeFieldVal("createcode_product_cate", "name", val["cid"])
			val["userType"] = gf.GetDicFieldVal("2", val["userType"])
		}
		gf.Success().SetMsg("获取全部列表").SetData(gf.Map{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}).Regin(c)
	}
}

// 保存
func (api *Product) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	var f_id = gf.GetEditId(param["id"])
	param["workerway"] = gf.ArrayToStr(param["workerway"], ",")
	if f_id == 0 {

		addId, err := gf.Model("createcode_product").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
		} else {
			if addId != 0 {
				gf.Model("createcode_product").Where("id", addId).Update(gf.Map{"weigh": addId})
			}
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("createcode_product").Where("id", f_id).Update(param)
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 更新状态
func (api *Product) UpStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res, err := gf.Model("createcode_product").Where("id", param["id"]).Update(param)
	if err != nil {
		gf.Failed().SetMsg("更新失败！").SetData(err).Regin(c)
	} else {
		msg := "更新成功！"
		if res == nil {
			msg = "暂无数据更新"
		}
		gf.Success().SetMsg(msg).SetData(res).Regin(c)
	}
}

// 删除
func (api *Product) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res, err := gf.Model("createcode_product").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res).Regin(c)
	}
}

// 获取内容
func (api *Product) GetContent(c *gf.GinCtx) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		gf.Failed().SetMsg("请传参数id").Regin(c)
	} else {
		data, err := gf.Model("createcode_product").Where("id", id).Find()
		if err != nil {
			gf.Failed().SetMsg("获取内容失败").SetData(err).Regin(c)
		} else {
			if data != nil && data["workerway"].String() != "" {
				data["workerway"] = gf.VarNew(gf.SplitAndStr(data["workerway"].String(), ","))
			}
			gf.Success().SetMsg("获取内容成功！").SetData(data).Regin(c)
		}
	}
}
