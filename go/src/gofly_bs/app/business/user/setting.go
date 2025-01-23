package user

import (
	"gofly/utils/gf"
	"gofly/utils/tools/gtime"
	"gofly/utils/tools/gvar"
	"strings"
)

// 修改账号信息
type Setting struct{}

func init() {
	fpath := Setting{}
	gf.Register(&fpath, fpath)
}
func (api *Setting) GetUserinfo(c *gf.GinCtx) {
	userID := c.GetInt64("userID") //当前用户ID
	userdata, err := gf.Model("business_account").Fields("id,username,dept_id,nickname,city,company,avatar,status,createtime").Where("id", userID).Find()
	if err != nil {
		gf.Failed().SetMsg("查找用户数据！").Regin(c)
	} else {
		rooturl := gf.GetMainURL()
		if userdata["avatar"] == nil {
			userdata["avatar"] = gvar.New(gf.GetMainURLLocal() + "resource/uploads/static/avatar.png")
		} else if !strings.Contains(userdata["avatar"].String(), "http") {
			userdata["avatar"] = gvar.New(rooturl + userdata["avatar"].String())
		}
		userdata["deptname"], _ = gf.Model("business_auth_dept").Where("id", userdata["dept_id"]).Value("name")
		roles, _ := gf.Model("business_auth_role_access").As("a").LeftJoin("business_auth_role", "r", "r.id=a.role_id").Where("a.uid", userID).Array("r.name")
		userdata["roles"] = gf.VarNew(strings.Join(gf.Strings(roles), "，"))
		gf.Success().SetMsg("获取用户信息").SetData(userdata).Regin(c)
	}
}

// 更新账号信息
func (api *Setting) SaveInfo(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	userID := c.GetInt64("userID") //当前用户ID
	if oldpassword, ok := param["oldpassword"]; ok {
		if param["type"] == "mobile" {
			code, emerr := gf.GetVerifyCode(gf.String(param["mobile"]))
			if emerr != nil || code != gf.Int(param["captcha"]) {
				gf.Failed().SetMsg("验证码无效").SetData(emerr).Regin(c)
				return
			}
			delete(param, "captcha")
		} else if param["type"] == "email" {
			code, emerr := gf.GetVerifyCode(gf.String(param["email"]))
			if emerr != nil || code != gf.Int(param["captcha"]) {
				gf.Failed().SetMsg("验证码无效").SetData(emerr).Regin(c)
				return
			}
			delete(param, "captcha")
		}
		account, err := gf.Model("business_account").Where("id", userID).Fields("password,salt").Find()
		if err != nil {
			gf.Failed().SetMsg("查找账号信息失败").SetData(err).Regin(c)
			return
		}
		salt := account["salt"].String()
		oldpass := gf.Md5(gf.String(oldpassword) + salt)
		if oldpass != account["password"].String() {
			gf.Failed().SetMsg("输入的当前密码不正确！").Regin(c)
			return
		}
		delete(param, "oldpassword")
		if param["type"] == "password" {
			param["password"] = gf.Md5(gf.String(param["newpassword"]) + salt)
			param["pwd_reset_time"] = gtime.Datetime()
			delete(param, "newpassword")
		}
	}
	res, err := gf.Model("business_account").Data(param).Where("id", userID).Update()
	if err != nil {
		gf.Failed().SetMsg("更新失败").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("更新成功！").SetData(res).Regin(c)
	}
}
