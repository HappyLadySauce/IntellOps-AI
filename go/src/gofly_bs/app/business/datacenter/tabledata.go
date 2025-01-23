package datacenter

import (
	"gofly/utils/gf"
)

type Tabledata struct{}

func init() {
	fpath := Tabledata{}
	gf.Register(&fpath, fpath)
}

// 获取列表
func (api *Tabledata) GetList(c *gf.GinCtx) {
	list, _ := gf.Model("common_dictionary_group").Fields("id,business_id,title,remark,tablename,status,weigh,data_from,db_way").Order("weigh asc").Select()
	gf.Success().SetMsg("获取列表").SetData(list).Regin(c)
}

// 保存
func (api *Tabledata) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	var f_id = gf.GetEditId(param["id"])
	if f_id == 0 {
		businessID, _ := c.Get("businessID") //当前用户businessID(saas账号ID)
		param["data_from"] = "business"
		param["business_id"] = businessID
		if param["db_way"] == "sys" {
			param["tablename"] = "common_dictionary_data"
		}
		addId, err := gf.Model("common_dictionary_group").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
		} else {
			if addId != 0 && gf.Int(param["weigh"]) == 0 {
				gf.Model("common_dictionary_group").Data(gf.Map{"weigh": addId}).Where("id", addId).Update()
			}
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("common_dictionary_group").Data(param).Where("id", f_id).Update()
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 删除
func (api *Tabledata) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	tablename, _ := gf.Model("common_dictionary_group").WhereIn("id", param["ids"]).Value("tablename")
	res2, err := gf.Model("common_dictionary_group").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Model(tablename.String()).WhereIn("group_id", param["ids"]).Delete()
		gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
	}
}
