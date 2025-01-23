package user

import (
	"gofly/utils/gf"
)

// 用于自动注册路由
type Account struct {
	NoNeedAuths []string
}

func init() {
	fpath := Account{NoNeedAuths: []string{"GetMenu"}}
	gf.Register(&fpath, fpath)
}

// 系统设置-获取菜单
func (api *Account) GetMenu(c *gf.GinCtx) {
	user, exists := gf.GetUserInfo(c) //当前用户
	if !exists {
		rouid := c.DefaultQuery("rouid", "")
		if rouid != "" {
			nemu_list, _ := gf.Model("business_auth_rule").Where("id", rouid).Order("weigh asc").Select()
			rulemenu := GetMenuArray(nemu_list, 0, make([]interface{}, 0))
			gf.Success().SetMsg("获取菜单").SetData(rulemenu).Regin(c)
			return
		}
		gf.Failed().SetMsg("登录失效").SetCode(401).Regin(c)
		return
	}
	//获取用户权限菜单
	role_id, acerr := gf.Model("business_auth_role_access").Where("uid", user.ID).Array("role_id")
	if acerr != nil {
		gf.Failed().SetMsg("获取role_access失败").SetData(acerr).Regin(c)
	}
	if role_id == nil {
		gf.Failed().SetMsg("您没有使用权限").Regin(c)
	}
	menu_ids, rerr := gf.Model("business_auth_role").WhereIn("id", role_id).Array("rules")
	if rerr != nil {
		gf.Failed().SetMsg("查找auth_role败！").SetData(rerr).Regin(c)
	}
	//获取超级角色
	super_role, _ := gf.Model("business_auth_role").WhereIn("id", role_id).Where("rules", "*").Value("id")
	RMDB := gf.Model("business_auth_rule")
	var roles []interface{}
	if super_role == nil { //不是超级权限-过滤菜单权限
		getmenus := gf.ArrayMerge(menu_ids)
		RMDB = RMDB.WhereIn("id", getmenus)
		roles = getmenus
	} else {
		roles = make([]interface{}, 0)
	}
	nemu_list, ruleerr := RMDB.Where("status", 0).WhereIn("type", []interface{}{0, 1}).Order("weigh asc").Select()
	if ruleerr != nil {
		gf.Failed().SetMsg("获取菜单错误").SetData(ruleerr).Regin(c)
		return
	}
	rulemenu := GetMenuArray(nemu_list, 0, roles)
	gf.Success().SetMsg("获取菜单").SetData(rulemenu).Regin(c)
}

// 保存数据
func (api *Account) Upuserinfo(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	userID := c.GetInt64("userID")
	res, err := gf.Model("business_account").Data(param).Where("id", userID).Update()
	if err != nil {
		gf.Failed().SetMsg("更新头像失败！").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("更新用户数据成功").SetData(res).Regin(c)
	}
}

// 更新头像
func (api *Account) Upavatar(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	userID := c.GetInt64("userID")
	res, err := gf.Model("business_account").Where("id", userID).Data(map[string]interface{}{"avatar": param["url"]}).Update()
	if err != nil {
		gf.Failed().SetMsg("更新头像失败！").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("更新头像成功").SetData(res).Regin(c)
	}
}

// 修改密码
func (api *Account) Changepwd(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	userID := c.GetInt64("userID")
	userdata, err := gf.Model("business_account").Where("id", userID).Fields("password,salt").Find()
	if err != nil {
		gf.Failed().SetMsg("查找账号失败！").SetData(err).Regin(c)
	} else {
		pass := gf.Md5(param["passwordOld"].(string) + userdata["salt"].String())
		if userdata["password"].String() != pass {
			gf.Failed().SetMsg("原来密码输入错误！").SetData(err).Regin(c)
		} else {
			newpass := gf.Md5(param["passwordNew"].(string) + userdata["salt"].String())
			res, err := gf.Model("business_account").
				Data(map[string]interface{}{"password": newpass}).Where("id", userID).Update()
			if err != nil {
				gf.Failed().SetMsg("修改密码失败").SetData(err).Regin(c)
			} else {
				gf.Success().SetMsg("修改密码成功！").SetData(res).Regin(c)
			}
		}
	}
}
