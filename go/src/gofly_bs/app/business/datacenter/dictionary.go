package datacenter

import (
	"encoding/json"
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gmap"
	"gofly/utils/tools/gvar"
	"strings"
)

// 字典数据
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
	whereMap := gmap.New()
	whereMap.Set("group_id", param["group_id"])
	if title, ok := param["title"]; ok && title != "" {
		whereMap.Set("keyname like ?", "%"+gconv.String(title)+"%")
	}
	if status, ok := param["status"]; ok && status != "" {
		whereMap.Set("status", status)
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	MDB := gf.Model(param["tablename"]).Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Page(pageNo, pageSize).Order("id asc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		rooturl := gf.GetMainURL()
		for _, val := range list {
			if _, ok := val["image"]; ok && val["image"].String() != "" && !strings.Contains(val["image"].String(), "http") {
				val["image"] = gvar.New(rooturl + val["image"].String())
			}
		}
		gf.Success().SetMsg("获取全部列表1").SetData(gf.Map{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}).Regin(c)
	}
}

// 保存
func (api *Dictionary) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	var f_id = gf.GetEditId(param["id"])
	tablename := param["tablename"]
	if f_id == 0 {
		param["data_from"] = "business"
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
	res, err := gf.Model(param["tablename"]).Where("id", param["id"]).Data(map[string]interface{}{"status": param["status"]}).Update()
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
func (api *Dictionary) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res, err := gf.Model(param["tablename"]).WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res).Regin(c)
	}
}

// 使用数据表数据-表单生成使用
func (api *Dictionary) GetTableDataForm(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	var custom interface{}
	if customstr, ok := param["custom"]; ok && customstr != "" {
		custom = gf.StringToJSON(gf.String(customstr))
	}
	if gf.DbHaseField(gf.String(param["tablename"]), "pid") {
		list, err := gf.Model(param["tablename"]).Where(custom).Fields("id,id as value,pid," + param["showfield"].(string) + " as label").Select()
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
		list, err := gf.Model(param["tablename"]).Where(custom).Fields("id as value," + param["showfield"].(string) + " as label").Select()
		if err != nil {
			gf.Failed().SetMsg("使用数据表数据失败！").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("使用数据表数据列表").SetData(list).Regin(c)
		}
	}
}
