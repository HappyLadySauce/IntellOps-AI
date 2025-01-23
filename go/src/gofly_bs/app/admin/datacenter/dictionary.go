package datacenter

import (
	"encoding/json"
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gmap"
)

// 字典数据管理
type Dictionary struct{ NoNeedAuths []string }

func init() {
	fpath := Dictionary{NoNeedAuths: []string{"getTableDataForm"}}
	gf.Register(&fpath, fpath)
}

// 获取列表
func (api *Dictionary) GetList(c *gf.GinCtx) {
	pageNo := gconv.Int(c.DefaultQuery("page", "1"))
	pageSize := gconv.Int(c.DefaultQuery("pageSize", "10"))
	//搜索添条件
	param, _ := gf.RequestParam(c)
	tablename := gconv.String(param["tablename"])
	whereMap := gmap.New()
	whereMap.Set("group_id", param["group_id"])
	if name, ok := param["name"]; ok && name != "" {
		whereMap.Set("name like ?", "%"+gconv.String(name)+"%")
	}
	if status, ok := param["status"]; ok && status != "" {
		whereMap.Set("status", status)
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	MDB := gf.Model(tablename).Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Page(pageNo, pageSize).Order("id asc").Select()
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

// 保存
func (api *Dictionary) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	tablename := param["tablename"]
	var f_id = gf.GetEditId(param["id"])
	if f_id == 0 {
		param["data_from"] = "common"
		addId, err := gf.Model(tablename).Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
		} else {
			if addId != 0 {
				gf.Model(tablename).Data(map[string]interface{}{"weigh": addId}).Where("id", addId).Update()
			}
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model(tablename).Data(param).Where("id", f_id).Update()
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 更新状态
func (api *Dictionary) UpStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model(param["tablename"]).Where("id", param["id"]).Data(map[string]interface{}{"status": param["status"]}).Update()
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
func (api *Dictionary) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model(param["tablename"]).WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
	}
}

// 使用数据表数据-表单生成使用
func (api *Dictionary) GetTableDataForm(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	if gf.DbHaseField(gf.String(param["tablename"]), "pid") {
		list, err := gf.Model(param["tablename"]).Fields("id,id as value,pid," + param["showfield"].(string) + " as label").Select()
		if err != nil {
			gf.Failed().SetMsg("使用数据表数据失败！").SetData(err).Regin(c)
		} else {
			list = gf.GetTreeArray(list, 0, "")
			dataTolist := list.Json()
			var parameter []gf.Map
			_ = json.Unmarshal([]byte(dataTolist), &parameter)
			listarray := gf.GetTreeToList(parameter, "label")
			gf.Success().SetMsg("使用数据表数据列表").SetData(listarray).Regin(c)
		}
	} else {
		list, err := gf.Model(param["tablename"]).Fields("id as value," + param["showfield"].(string) + " as label").Select()
		if err != nil {
			gf.Failed().SetMsg("使用数据表数据失败！").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("使用数据表数据列表").SetData(list).Regin(c)
		}
	}
}
