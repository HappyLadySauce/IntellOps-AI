package system

import (
	"fmt"
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gmap"
	"gofly/utils/tools/grand"
	"gofly/utils/tools/gvar"
)

// 账号管理
type Account struct{ NoNeedAuths []string }

func init() {
	fpath := Account{NoNeedAuths: []string{"GetRole", "isaccountexist"}}
	gf.Register(&fpath, fpath)
}

// 获取成员列表
func (api *Account) GetList(c *gf.GinCtx) {
	pageNo := gconv.Int(c.DefaultQuery("page", "1"))
	pageSize := gconv.Int(c.DefaultQuery("pageSize", "10"))
	//搜索添条件
	param, _ := gf.RequestParam(c)
	whereMap := gmap.New()
	account_id, filter := gf.GetDataAuthor(c)
	if filter { //需要权限过滤
		whereMap.Set("account_id IN(?)", account_id)
	}
	if cid, ok := param["cid"]; ok && cid != "0" {
		whereMap.Set("dept_id", cid)
	}
	if name, ok := param["name"]; ok && name != "" {
		whereMap.Set("name like ?", "%"+gconv.String(name)+"%")
	}
	if mobile, ok := param["mobile"]; ok && mobile != "" {
		whereMap.Set("mobile like ?", "%"+gconv.String(mobile)+"%")
	}
	if status, ok := param["status"]; ok && status != "" {
		whereMap.Set("status", status)
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	MDB := gf.Model("admin_account").Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Fields("id,status,name,username,avatar,tel,mobile,email,remark,dept_id,remark,city,address,company,createtime").Page(pageNo, pageSize).Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		for _, val := range list {
			roleid, _ := gf.Model("admin_auth_role_access").Where("uid", val["id"]).Array("role_id")
			rolename, _ := gf.Model("admin_auth_role").WhereIn("id", roleid).Array("name")
			val["rolename"] = gvar.New(rolename)
			val["roleid"] = gvar.New(roleid)
			depname, _ := gf.Model("admin_auth_dept").Where("id", val["dept_id"]).Value("name")
			val["depname"] = depname
			//头像
			if val["avatar"] == nil {
				val["avatar"] = gvar.New("resource/uploads/static/avatar.png")
			}
		}
		gf.Success().SetMsg("获取全部列表").SetData(gf.Map{
			"page":     pageNo,
			"pageSize": pageSize,
			"total":    totalCount,
			"items":    list}).Regin(c)
	}
}

// 保存、编辑
func (api *Account) Save(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	var roleid []interface{}
	if param["roleid"] != nil {
		roleid = param["roleid"].([]interface{})
		delete(param, "roleid")
	}
	if param["password"] != nil && param["password"] != "" {
		salt := grand.Str("123456789", 6)
		mdpass := fmt.Sprintf("%v%v", param["password"], salt)
		param["password"] = gf.Md5(mdpass)
		param["salt"] = salt
	}
	if param["avatar"] == "" {
		param["avatar"] = "resource/uploads/static/avatar.png"
	}
	var f_id = gf.GetEditId(param["id"])
	if f_id == 0 {
		param["account_id"] = c.GetInt64("userID")
		addId, err := gf.Model("admin_account").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
		} else {
			//添加角色-多个
			appRoleAccess(roleid, addId)
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("admin_account").
			Data(param).
			Where("id", f_id).
			Update()
		if err != nil {
			gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
		} else {
			//添加角色-多个
			if roleid != nil {
				appRoleAccess(roleid, f_id)
			}
			gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
		}
	}
}

// 更新状态
func (api *Account) UpStatus(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("admin_account").Where("id", param["id"]).Data(map[string]interface{}{"status": param["status"]}).Update()
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
func (api *Account) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res2, err := gf.Model("admin_account").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res2).Regin(c)
	}
}

// 获取内容
func (api *Account) GetContent(c *gf.GinCtx) {
	id := c.DefaultQuery("id", "")
	if id == "" {
		gf.Failed().SetMsg("请传参数id").Regin(c)
	} else {
		data, err := gf.Model("admin_account").Where("id", id).Find()
		if err != nil {
			gf.Failed().SetMsg("获取内容失败").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("获取内容成功！").SetData(data).Regin(c)
		}
	}
}

// 添加授权
func appRoleAccess(roleids []interface{}, uid interface{}) {
	//批量提交
	gf.Model("admin_auth_role_access").Where("uid", uid).Delete()
	save_arr := []map[string]interface{}{}
	for _, val := range roleids {
		marr := map[string]interface{}{"uid": uid, "role_id": val}
		save_arr = append(save_arr, marr)
	}
	gf.Model("admin_auth_role_access").Data(save_arr).Insert()
}

// 表单-选择角色
func (api *Account) GetRole(c *gf.GinCtx) {
	userID := c.GetInt64("userID") //当前用户ID
	user_role_ids, _ := gf.Model("admin_auth_role_access").Where("uid", userID).Array("role_id")
	role_ids := gf.GetAllChilIds("admin_auth_role", user_role_ids) //批量获取子节点id
	all_role_id := gf.MergeArr(user_role_ids, role_ids)
	menuList, _ := gf.Model("admin_auth_role").WhereIn("id", all_role_id).Where("status", 0).Fields("id ,pid,name").Order("weigh asc").Select()
	//获取最大一级的pid
	max_role_id, _ := gf.Model("admin_auth_role").Where("id", user_role_ids).Order("id asc").Value("pid")
	list := gf.GetMenuChildrenArray(menuList, gf.Int64(max_role_id), "pid")
	gf.Success().SetMsg("表单选择角色多选用数据").SetData(list).Regin(c)
}

// 判断账号是否存在
func (api *Account) Isaccountexist(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	if param["id"] != nil {
		res1, err := gf.Model("admin_account").Where("id !=?", param["id"]).Where("username", param["username"]).Value("id")
		if err != nil {
			gf.Failed().SetMsg("验证失败").SetData(err).Regin(c)
		} else if res1 != nil {
			gf.Failed().SetMsg("账号已存在").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("验证通过").SetData(res1).Regin(c)
		}
	} else {
		res2, err := gf.Model("admin_account").Where("username", param["username"]).Value("id")
		if err != nil {
			gf.Failed().SetMsg("验证失败").SetData(err).Regin(c)
		} else if res2 != nil {
			gf.Failed().SetMsg("账号已存在").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("验证通过").SetData(res2).Regin(c)
		}
	}
}
