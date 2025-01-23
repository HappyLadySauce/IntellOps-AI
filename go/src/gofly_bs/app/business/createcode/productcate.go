package createcode

import (
	"gofly/utils/gf"
	"gofly/utils/tools/gmap"
)

// 关联的分类
type Productcate struct{ NoNeedAuths []string }

func init() {
	fpath := Productcate{NoNeedAuths: []string{"getTree"}}
	gf.Register(&fpath, fpath)
}

// 获取列表-tree
func (api *Productcate) GetTree(c *gf.GinCtx) {
	//搜索添条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	whereMap.Set("business_id", c.GetInt64("businessID"))
	if name, ok := param["name"]; ok && name != "" {
		whereMap.Set("name like ?", "%"+gf.String(name)+"%")
	}
	if status, ok := param["status"]; ok && status != "" {
		whereMap.Set("status", status)
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	list, err := gf.Model("createcode_product_cate").Where(whereMap).Order("id asc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		for _, val := range list {
			val["key"] = val["id"]
		}
		list = gf.GetMenuChildrenArray(list, 0, "pid")
		gf.Success().SetMsg("获取树形列表").SetData(list).Regin(c)
	}
}

// 保存
func (api *Productcate) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	var f_id = gf.GetEditId(param["id"])
	if f_id == 0 {
		param["business_id"] = c.GetInt64("businessID") //当前用户商户ID
		addId, err := gf.Model("createcode_product_cate").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
		} else {
			if addId != 0 {
				gf.Model("createcode_product_cate").Where("id", addId).Update(gf.Map{"weigh": addId})
			}
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("createcode_product_cate").Where("business_id", c.GetInt64("businessID")).Where("id", f_id).Update(param)
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 更新状态
func (api *Productcate) UpStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res, err := gf.Model("createcode_product_cate").Where("business_id", c.GetInt64("businessID")).Where("id", param["id"]).Update(param)
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
func (api *Productcate) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res, err := gf.Model("createcode_product_cate").Where("business_id", c.GetInt64("businessID")).WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Model("createcode_product_cate").Where("business_id", c.GetInt64("businessID")).WhereIn("pid", param["ids"]).Delete()
		gf.Success().SetMsg("删除成功！").SetData(res).Regin(c)
	}
}
