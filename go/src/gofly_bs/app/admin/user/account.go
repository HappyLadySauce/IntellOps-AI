package user

import (
	"encoding/json"
	"gofly/utils/gf"
	"io"
)

type Account struct {
	NoNeedAuths []string
}

func init() {
	fpath := Account{NoNeedAuths: []string{"GetMenu"}}
	gf.Register(&fpath, fpath)
}

// 系统设置-获取菜单
func (api *Account) GetMenu(c *gf.GinCtx) {
	userID := c.GetInt64("userID") //当前用户ID
	//获取用户权限菜单
	role_id, acerr := gf.Model("admin_auth_role_access").Where("uid", userID).Array("role_id")
	if acerr != nil {
		gf.Failed().SetMsg("获取role_access失败").SetData(acerr).Regin(c)
	}
	if role_id == nil {
		gf.Failed().SetMsg("您没有使用权限").Regin(c)
	}
	menu_ids, rerr := gf.Model("admin_auth_role").WhereIn("id", role_id).Array("rules")
	if rerr != nil {
		gf.Failed().SetMsg("查找auth_role败！").SetData(acerr).Regin(c)
	}
	//获取超级角色
	super_role, _ := gf.Model("admin_auth_role").WhereIn("id", role_id).Where("rules", "*").Value("id")
	RMDB := gf.Model("admin_auth_rule")
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
	}
	rulemenu := GetMenuArray(nemu_list, 0, roles)
	gf.Success().SetMsg("获取菜单").SetData(rulemenu).Regin(c)
}

// 保存数据
func (api *Account) Upuserinfo(c *gf.GinCtx) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	userID := c.GetInt64("userID")
	res, err := gf.Model("admin_account").
		Data(parameter).
		Where("id", userID).
		Update()
	if err != nil {
		gf.Success().SetMsg("更新失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("更新用户数据成功").SetData(res).Regin(c)
	}
}

// 更新头像
func (api *Account) Upavatar(c *gf.GinCtx) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	userID := c.GetInt64("userID")
	res, err := gf.Model("admin_account").Where("id", userID).Data(map[string]interface{}{"avatar": parameter["url"]}).Update()
	if err != nil {
		gf.Failed().SetMsg("更新头像失败！").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("更新头像成功").SetData(res).Regin(c)
	}
}

// 修改密码
func (api *Account) Changepwd(c *gf.GinCtx) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	_ = json.Unmarshal(body, &parameter)
	userID := c.GetInt64("userID")
	userdata, err := gf.Model("admin_account").Where("id", userID).Fields("password,salt").Find()
	if err != nil {
		gf.Failed().SetMsg("查找账号失败！").SetData(err).Regin(c)
	} else {
		pass := gf.Md5(parameter["passwordOld"].(string) + userdata["salt"].String())
		if userdata["password"].String() != pass {
			gf.Failed().SetMsg("原来密码输入错误！").SetData(err).Regin(c)
		} else {
			newpass := gf.Md5(parameter["passwordNew"].(string) + userdata["salt"].String())
			res, err := gf.Model("admin_account").
				Data(map[string]interface{}{"password": newpass}).
				Where("id", userID).
				Update()
			if err != nil {
				gf.Failed().SetMsg("修改密码失败").SetData(err).Regin(c)
			} else {
				gf.Success().SetMsg("修改密码成功！").SetData(res).Regin(c)
			}
		}
	}
}
