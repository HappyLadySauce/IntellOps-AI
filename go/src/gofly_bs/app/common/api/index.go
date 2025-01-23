package api

import (
	"gofly/utils/gf"
	"gofly/utils/tools/gmap"
)

// 低代码api-通用
// 内置常用的curd、内置公共接口
type Index struct{}

func init() {
	fpath := Index{}
	gf.Register(&fpath, fpath)
}

// 1.获取数据列表
func (api *Index) GetList(c *gf.GinCtx) {
	api_id := c.DefaultQuery("api_id", "0")
	if gf.Int(api_id) == 0 {
		gf.Failed().SetMsg("请传接口api_id").Regin(c)
		return
	}
	apidata, _ := gf.Model("common_api").Where("id", api_id).Find()
	pageNo := gf.Int(c.DefaultQuery("page", "1"))
	pageSize := gf.Int(c.DefaultQuery("pageSize", "10"))
	//搜索添条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	if gf.DbHaseField(apidata["tablename"].String(), "business_id") {
		businessID, _ := c.Get("businessID") //当前用户ID
		whereMap.Set("business_id", businessID)
	}
	if title, ok := param["title"]; ok && title != "" {
		whereMap.Set("title like ?", "%"+gf.String(title)+"%")
	}
	if status, ok := param["status"]; ok && status != "" {
		whereMap.Set("status", status)
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	MDB := gf.Model(apidata["tablename"].String()).Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Fields(apidata["fields"].String()).Page(pageNo, pageSize).Order("id desc").Select()
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

// 2.获取单条数据
func (api *Index) GetData(c *gf.GinCtx) {
	api_id := c.DefaultQuery("api_id", "0")
	if gf.Int(api_id) == 0 {
		gf.Failed().SetMsg("请传接口api_id").Regin(c)
		return
	}
	apidata, _ := gf.Model("common_api").Where("id", api_id).Find()
	id := c.DefaultQuery("id", "")
	if id == "" {
		gf.Failed().SetMsg("请传参数id").Regin(c)
	} else {
		data, err := gf.Model(apidata["tablename"].String()).Where("id", id).Find()
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

// 3.添加数据
func (api *Index) Add(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	if api_id, ok := param["api_id"]; !ok || gf.Int(api_id) == 0 {
		gf.Failed().SetMsg("请传接口api_id").Regin(c)
		return
	}
	apidata, _ := gf.Model("common_api").Where("id", param["api_id"]).Find()
	if gf.DbHaseField(apidata["tablename"].String(), "business_id") {
		param["business_id"], _ = c.Get("businessID") //当前用户ID
	}
	addId, err := gf.Model(apidata["tablename"].String()).Data(param).InsertAndGetId()
	if err != nil {
		gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
	} else {
		if addId != 0 {
			gf.Model(apidata["tablename"].String()).Data(gf.Map{"weigh": addId}).Where("id", addId).Update()
		}
		gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
	}
}

// 4.更新数据
func (api *Index) Put(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	if api_id, ok := param["api_id"]; !ok || gf.Int(api_id) == 0 {
		gf.Failed().SetMsg("请传接口api_id").Regin(c)
		return
	}
	apidata, _ := gf.Model("common_api").Where("id", param["api_id"]).Find()
	if id, ok := param["id"]; !ok || gf.Int(id) == 0 {
		gf.Failed().SetMsg("请传修改的数据id").Regin(c)
	} else {
		res, err := gf.Model(apidata["tablename"].String()).Data(param).Where("id", param["id"]).Update()
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 5.删除数据
func (api *Index) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	if api_id, ok := param["api_id"]; !ok || gf.Int(api_id) == 0 {
		gf.Failed().SetMsg("请传接口api_id").Regin(c)
		return
	}
	apidata, _ := gf.Model("common_api").Where("id", param["api_id"]).Find()
	_, err := gf.Model(apidata["tablename"].String()).WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").Regin(c)
	}
}
