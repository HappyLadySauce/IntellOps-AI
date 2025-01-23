package business

import (
	"fmt"
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gmap"
	"gofly/utils/tools/gvar"
	"strings"
)

type Bizrole struct{ NoNeedAuths []string }

func init() {
	fpath := Bizrole{NoNeedAuths: []string{"getParent", "getMenuList"}}
	gf.Register(&fpath, fpath)
}

// 获取数据列表-子树结构
func (api *Bizrole) GetList(c *gf.GinCtx) {
	userID := c.GetInt64("userID") //当前用户ID
	role_id, _ := gf.Model("business_auth_role_access").Where("uid", userID).Array("role_id")
	role_ids := gf.GetAllChilIds("business_auth_role", role_id) //批量获取子节点id
	all_role_id := gf.MergeArr(role_id, role_ids)
	//查找条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	whereMap.Set("id", all_role_id) //in 查询
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
	roleList, _ := gf.Model("business_auth_role").Where(whereMap).Order("weigh asc").Select()
	roleList = gf.GetTreeArray(roleList, 0, "")
	if roleList == nil {
		roleList = make(gf.OrmResult, 0)
	}
	gf.Success().SetMsg("获取拥有角色列表").SetData(roleList).Regin(c)
}

// 表单获取选择父级
func (api *Bizrole) GetParent(c *gf.GinCtx) {
	id := c.DefaultQuery("id", "0")
	menuList, _ := gf.Model("business_auth_role").Where("id !=", id).Fields("id,pid,name").Order("weigh asc").Select()
	max_role_id, _ := gf.Model("business_auth_role").Where("id !=", id).Order("id asc").Value("pid")
	menuList = gf.GetMenuChildrenArray(menuList, gf.Int64(max_role_id), "pid")
	gf.Success().SetMsg("部门父级数据！").SetData(menuList).Regin(c)
}

// 表单获取菜单
func (api *Bizrole) GetMenuList(c *gf.GinCtx) {
	pid := c.DefaultQuery("pid", "0")
	MDB := gf.Model("business_auth_rule").Where("status", 0).WhereIn("type", []interface{}{0, 1})
	var rule_ids []interface{}
	if pid == "0" { //获取本账号所拥有的权限
		userID := c.GetInt64("userID")
		role_id, _ := gf.Model("business_auth_role_access").Where("uid", userID).Array("role_id")
		menu_id, _ := gf.Model("business_auth_role").WhereIn("id", role_id).Array("rules")
		super_role, _ := gf.Model("business_auth_role").WhereIn("id", role_id).Where("rules", "*").Value("id")
		if super_role == nil { //不是超级权限-过滤菜单权限
			getmenus := gf.ArraymoreMerge(menu_id)
			MDB = MDB.WhereIn("id", getmenus)
			rule_ids = getmenus
		}
	} else {
		//获取用户权限
		menu_id_str, _ := gf.Model("business_auth_role").Where("id", pid).Value("rules")
		if !strings.Contains(menu_id_str.String(), "*") { //不是超级权限-过滤菜单权限
			getmenus := gf.Axplode(menu_id_str.String())
			MDB = MDB.WhereIn("id", getmenus)
			rule_ids = getmenus
		}
	}
	menuList, _ := MDB.Fields("id,pid,title,locale").Order("weigh asc").Select()
	for _, val := range menuList {
		if val["title"].String() == "" {
			val["title"] = val["locale"]
		}
		delete(val, "locale")
		//获取按钮
		whereMap := gmap.New()
		if rule_ids != nil {
			whereMap.Set("id IN(?)", rule_ids)
		}
		btn_rules, _ := gf.Model("business_auth_rule").Where("status", 0).Where("type", 2).Where("pid", val["id"]).Where(whereMap).Fields("id,pid,title,locale").Order("weigh asc").Select()
		if btn_rules != nil && len(btn_rules) > 0 {
			item := gf.Map{
				"title":     "按钮权限",
				"id":        btn_rules[0]["id"],
				"pid":       val["id"],
				"checkable": false,
				"btn_rules": btn_rules,
			}
			var valitem []gf.Map
			valitem = append(valitem, item)
			val["children"] = gvar.New(valitem)
			var btnids []interface{}
			for _, btnid := range btn_rules {
				btnids = append(btnids, btnid["id"])
			}
			val["btnids"] = gvar.New(btnids)
		} else if val["pid"].Int() == 0 {
			//一级菜单获取子级菜单按钮
			sub_rule_ids, _ := gf.Model("business_auth_rule").Where("pid", val["id"]).Where("status", 0).Where("type !=", 2).Array("id")
			btn_rule_ids, _ := gf.Model("business_auth_rule").Where("status", 0).Where("type", 2).WhereIn("pid", sub_rule_ids).Array("id")
			val["btnids"] = gvar.New(btn_rule_ids)
		}
		val["checkable"] = gvar.New(true)
	}
	menuList = gf.GetMenuChildrenArray(menuList, 0, "pid")
	if rule_ids == nil {
		btn_idsdata, _ := gf.Model("business_auth_rule").Where("status", 0).Where("type", 2).Array("id")
		gf.Success().SetMsg("获取菜单数据1").SetData(gf.Map{"list": menuList, "btn_rule_ids": btn_idsdata}).Regin(c)
	} else {
		btn_idsdata, _ := gf.Model("business_auth_rule").Where("status", 0).Where("type", 2).WhereIn("id", rule_ids).Array("id")
		gf.Success().SetMsg("获取菜单数据2").SetData(gf.Map{"list": menuList, "btn_rule_ids": btn_idsdata}).Regin(c)
	}
}

// 保存编辑
func (api *Bizrole) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	var f_id = gf.GetEditId(param["id"])
	if param["menu"] != nil && param["menu"] != "*" {
		rules := gf.GetRulesID("business_auth_rule", "pid", param["menu"]) //获取子菜单包含的父级ID
		rudata := rules.([]interface{})
		var rulesStr []string
		for _, v := range rudata {
			str := fmt.Sprintf("%v", v) //interface{}强转string
			rulesStr = append(rulesStr, str)
		}
		for _, bv := range param["btns"].([]interface{}) {
			str := fmt.Sprintf("%v", bv) //interface{}强转string
			rulesStr = append(rulesStr, str)
		}
		param["rules"] = strings.Join(rulesStr, ",")
	}
	if f_id == 0 {
		addId, err := gf.Model("business_auth_role").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
		} else {
			if addId != 0 {
				gf.Model("business_auth_role").Data(map[string]interface{}{"weigh": addId}).Where("id", addId).Update()
			}
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("business_auth_role").Data(param).Where("id", f_id).Update()
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 更新状态
func (api *Bizrole) UpStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("business_auth_role").Where("id", param["id"]).Data(map[string]interface{}{"status": param["status"]}).Update()
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
func (api *Bizrole) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("business_auth_role").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
	}
}
