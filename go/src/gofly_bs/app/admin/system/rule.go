package system

import (
	"gofly/utils/gf"
	"strings"
)

// 菜单管理
type Rule struct{ NoNeedAuths []string }

func init() {
	fpath := Rule{NoNeedAuths: []string{"getParent", "getRoutes", "getContent"}}
	gf.Register(&fpath, fpath)
}

// 1获取列表
func (api *Rule) GetList(c *gf.GinCtx) {
	menuList, _ := gf.Model("admin_auth_rule").Fields("id,pid,type,title,locale,icon,permission,path,component,weigh,status,createtime").Order("weigh asc").Select()
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

// 2获取列表-获取选项列表
func (api *Rule) GetParent(c *gf.GinCtx) {
	id := c.DefaultQuery("id", "0")
	menuList, err := gf.Model("admin_auth_rule").WhereIn("type", []interface{}{0, 1}).Where("id !=", id).Fields("id,pid,title,locale,routepath").Order("weigh asc").Select()
	if err != nil {
		gf.Failed().SetMsg("获取数据失败").SetData(err).Regin(c)
	} else {
		if menuList == nil {
			menuList = make(gf.OrmResult, 0)
		}
		for _, val := range menuList {
			if val["title"].String() == "" {
				val["title"] = val["locale"]
			}
		}
		menuTree := gf.GetMenuChildrenArray(menuList, 0, "pid")
		gf.Success().SetMsg("菜单父级数据！").SetData(gf.Map{"tree": menuTree, "list": menuList}).Regin(c)
	}
}

// 获取权限选择的路由列表
func (api *Rule) GetRoutes(c *gf.GinCtx) {
	filePath := "runtime/app/routers.txt"
	list := gf.ReaderFileByline(filePath)
	var nlist []interface{} = make([]interface{}, 0)
	for _, val := range list {
		item := strings.Split(gf.String(val), ":")
		if len(item) == 2 && !strings.HasPrefix(item[1], "/business/") && !strings.HasPrefix(item[1], "/common/") {
			nlist = append(nlist, gf.Map{"method": item[0], "path": item[1]})
		}
	}
	gf.Success().SetMsg("获取权限选择的路由列表").SetData(nlist).Regin(c)
}

// 3保存、编辑菜单
func (api *Rule) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	param["uid"] = c.GetInt64("userID") //当前用户ID
	var f_id = gf.GetEditId(param["id"])
	if f_id == 0 {
		addId, err := gf.Model("admin_auth_rule").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加菜单失败").SetData(err).Regin(c)
		} else {
			if addId != 0 {
				gf.Model("admin_auth_rule").Data(map[string]interface{}{"weigh": addId}).Where("id", addId).Update()
			}
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("admin_auth_rule").Data(param).Where("id", f_id).Update()
		if err != nil {
			gf.Failed().SetMsg("更新菜单失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 4更新状态
func (api *Rule) UpStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("admin_auth_rule").Where("id", param["id"]).Data(map[string]interface{}{"status": param["status"]}).Update()
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
func (api *Rule) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("admin_auth_rule").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除菜单失败").SetData(err).Regin(c)
	} else {
		//删除子类数据
		gf.Model("admin_auth_rule").WhereIn("pid", param["ids"]).Delete()
		gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
	}
}

// 获取内容
func (api *Rule) GetContent(c *gf.GinCtx) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		gf.Failed().SetMsg("请传参数id").Regin(c)
	} else {
		data, err := gf.Model("admin_auth_rule").Where("id", id).Find()
		if err != nil {
			gf.Failed().SetMsg("获取内容失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("获取内容成功！").SetData(data).Regin(c)
		}
	}
}
