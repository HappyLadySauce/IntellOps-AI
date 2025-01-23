package business

import (
	"gofly/utils/gf"
)

type Bizrule struct{ NoNeedAuths []string }

func init() {
	fpath := Bizrule{NoNeedAuths: []string{"getParent", "getRoutes", "getContent"}}
	gf.Register(&fpath, fpath)
}

// 1获取列表
func (api *Bizrule) GetList(c *gf.GinCtx) {
	menuList, _ := gf.Model("business_auth_rule").Fields("id,pid,type,title,locale,icon,permission,path,component,weigh,status,createtime").Order("weigh asc").Select()
	if menuList == nil {
		menuList = make(gf.OrmResult, 0)
	}
	for _, val := range menuList {
		if val["title"].String() == "" {
			val["title"] = val["locale"]
		}
	}
	menuList = gf.GetTreeArray(menuList, 0, "")
	gf.Success().SetMsg("获取全部菜单列表").SetData(menuList).Regin(c)
}

// 3保存、编辑菜单
func (api *Bizrule) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	param["uid"] = c.GetInt64("userID") //当前用户ID
	var f_id = gf.GetEditId(param["id"])
	if f_id == 0 {
		addId, err := gf.Model("business_auth_rule").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加菜单失败").SetData(err).Regin(c)
		} else {
			if addId != 0 {
				gf.Model("business_auth_rule").Data(map[string]interface{}{"weigh": addId}).Where("id", addId).Update()
			}
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("business_auth_rule").Data(param).Where("id", f_id).Update()
		if err != nil {
			gf.Failed().SetMsg("更新菜单失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 4更新状态
func (api *Bizrule) UpStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("business_auth_rule").Where("id", param["id"]).Data(map[string]interface{}{"status": param["status"]}).Update()
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

// 删除菜单
func (api *Bizrule) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("business_auth_rule").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除菜单失败").SetData(err).Regin(c)
	} else {
		//删除子类数据
		gf.Model("business_auth_rule").WhereIn("pid", param["ids"]).Delete()
		gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
	}
}

// 获取内容
func (api *Bizrule) GetContent(c *gf.GinCtx) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		gf.Failed().SetMsg("请传参数id").Regin(c)
	} else {
		data, err := gf.Model("business_auth_rule").Where("id", id).Find()
		if err != nil {
			gf.Failed().SetMsg("获取内容失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("获取内容成功！").SetData(data).Regin(c)
		}
	}
}
