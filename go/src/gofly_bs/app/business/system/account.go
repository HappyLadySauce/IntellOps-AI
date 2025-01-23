package system

//系统账号管理
import (
	"fmt"
	"gofly/utils/gf"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gmap"
	"gofly/utils/tools/grand"
	"gofly/utils/tools/gvar"
)

// 用户账号管理
type Account struct{ NoNeedAuths []string }

func init() {
	fpath := Account{NoNeedAuths: []string{"getList", "GetRole", "isaccountexist"}}
	gf.Register(&fpath, fpath)
}

// 获取成员列表
func (api *Account) GetList(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	businessID := c.GetInt64("businessID") //当前商户ID
	page := c.DefaultQuery("page", "1")
	_pageSize := c.DefaultQuery("pageSize", "10")
	pageNo := gconv.Int(page)
	pageSize := gconv.Int(_pageSize)
	//搜索添条件
	whereMap := gmap.New()
	whereMap.Set("business_id", businessID)
	account_id, filter := gf.GetDataAuthor(c)
	if filter { //需要权限过滤
		whereMap.Set("account_id IN(?)", account_id)
	}
	if name, ok := param["name"]; ok && name != "" {
		whereMap.Set("name like ?", "%"+gconv.String(name)+"%")
	}
	if mobile, ok := param["mobile"]; ok && mobile != "" {
		whereMap.Set("name", mobile)
	}
	if cid, ok := param["cid"]; ok && cid != "" {
		whereMap.Set("cid", cid)
	}
	if createtime, ok := param["createtime"]; ok && createtime != "" {
		datetime_arr := gf.SplitAndStr(gf.String(createtime), ",")
		whereMap.Set("createtime between ? and ?", gf.Slice{datetime_arr[0] + " 00:00", datetime_arr[1] + " 23:59"})
	}
	MDB := gf.Model("business_account").Where(whereMap)
	totalCount, _ := MDB.Clone().Count()
	list, err := MDB.Fields("id,status,name,username,avatar,tel,mobile,email,remark,dept_id,remark,city,address,company,createtime").
		Page(pageNo, pageSize).Order("id desc").Select()
	if err != nil {
		gf.Failed().SetMsg(err.Error()).Regin(c)
	} else {
		rooturl := gf.GetMainURL()
		for _, val := range list {
			roleid, _ := gf.Model("business_auth_role_access").Where("uid", val["id"]).Array("role_id")
			rolename, _ := gf.Model("business_auth_role").WhereIn("id", roleid).Array("name")
			val["rolename"] = gvar.New(rolename)
			val["roleid"] = gvar.New(roleid)
			depname, _ := gf.Model("business_auth_dept").Where("id", val["dept_id"]).Value("name")
			val["depname"] = depname
			//头像
			if val["avatar"] == nil {
				val["avatar"] = gvar.New(rooturl + "resource/uploads/static/avatar.png")
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
	var f_id = gf.GetEditId(param["id"])
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
	if f_id == 0 {
		param["account_id"] = c.GetInt64("userID")      //当前用户ID
		param["business_id"] = c.GetInt64("businessID") //当前商户ID
		addId, err := gf.Model("business_account").Data(param).InsertAndGetId()
		if err != nil {
			gf.Failed().SetMsg("添加失败").SetData(err).Regin(c)
		} else {
			//添加角色-多个
			appRoleAccess(roleid, addId)
			gf.Success().SetMsg("添加成功！").SetData(addId).Regin(c)
		}
	} else {
		res, err := gf.Model("business_account").Data(param).Where("id", f_id).Update()
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
	res, err := gf.Model("business_account").Where("id", param["id"]).Data(gf.Map{"status": param["status"]}).Update()
	if err != nil {
		gf.Failed().SetMsg("更新失败！").SetData(err).Regin(c)
	} else {
		msg := "更新成功！"
		if res != nil {
			msg = "暂无数据更新"
		}
		gf.Success().SetMsg(msg).Regin(c)
	}
}

// 删除
func (api *Account) Del(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	res, err := gf.Model("business_account").WhereIn("id", param["ids"]).Delete()
	if err != nil {
		gf.Failed().SetMsg("删除失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("删除成功！").SetData(res).Regin(c)
	}
}

// 添加授权
func appRoleAccess(roleids []interface{}, uid interface{}) {
	//批量提交
	gf.Model("business_auth_role_access").Where("uid", uid).Delete()
	save_arr := []map[string]interface{}{}
	for _, val := range roleids {
		marr := map[string]interface{}{"uid": uid, "role_id": val}
		save_arr = append(save_arr, marr)
	}
	gf.Model("business_auth_role_access").Data(save_arr).Save()
}

// 表单-选择角色
func (api *Account) GetRole(c *gf.GinCtx) {
	userID := c.GetInt64("userID") //当前用户ID
	user_role_ids, _ := gf.Model("business_auth_role_access").Where("uid", userID).Array("role_id")
	role_chil_ids := gf.GetAllChilIds("business_auth_role", user_role_ids) //批量获取子节点id
	all_role_id := gf.MergeArr(user_role_ids, role_chil_ids)
	//查找条件
	whereMap := gmap.New()
	whereMap.Set("id IN(?)", all_role_id) //in 查询
	account_id, _ := gf.GetDataAuthor(c)
	account_id = append(account_id, 0)
	//获取自己权限组-显示自己所在的权限组
	my_role_account_id, _ := gf.Model("business_auth_role").WhereIn("id", user_role_ids).Value("account_id")
	account_id = append(account_id, my_role_account_id)
	whereMap.Set("account_id IN(?)", account_id)
	roleList, _ := gf.Model("business_auth_role").Where(whereMap).Order("weigh asc").Select()
	//获取最大一级的pid
	max_role_id, _ := gf.Model("business_auth_role").Where(whereMap).Order("id asc").Value("pid")
	roleList = gf.GetTreeArray(roleList, gf.Int64(max_role_id), "")
	if roleList == nil {
		roleList = make(gf.OrmResult, 0)
	}
	gf.Success().SetMsg("表单选择角色多选用数据").SetData(roleList).Regin(c)
}

// 判断账号是否存在
func (api *Account) Isaccountexist(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	if param["id"] != nil {
		res1, err := gf.Model("business_account").Where("id !=", param["id"]).Where("username", param["username"]).Value("id")
		if err != nil {
			gf.Failed().SetMsg("验证失败").SetData(err).Regin(c)
		} else if res1 != nil {
			gf.Failed().SetMsg("账号已存在").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("验证通过").SetData(res1).Regin(c)
		}
	} else {
		res2, err := gf.Model("business_account").Where("username", param["username"]).Value("id")
		if err != nil {
			gf.Failed().SetMsg("验证失败").SetData(err).Regin(c)
		} else if res2 != nil {
			gf.Failed().SetMsg("账号已存在").SetData(err).Regin(c)
		} else {
			gf.Success().SetMsg("验证通过").SetData(res2).Regin(c)
		}
	}
}
