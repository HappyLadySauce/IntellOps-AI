package user

import (
	"gofly/utils/gf"
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gtime"
	"gofly/utils/tools/gvar"
	"strings"
)

/**
*使用 Index 是省略路径中的index
*本路径为： /business/user/login -省去了index
 */

type Index struct {
	NoNeedLogin []string //忽略登录接口配置-忽略全部传[*]
	NoNeedAuths []string //忽略RBAC权限认证接口配置-忽略全部传[*]
}

// 初始化路由
func init() {
	fpath := Index{NoNeedLogin: []string{"login", "logout"}, NoNeedAuths: []string{"*"}}
	gf.Register(&fpath, fpath)
}

/**
*1.《登录》
 */
func (api *Index) Login(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	if _, ok := param["username"]; ok {
		if param["username"] == nil || param["password"] == nil {
			gf.Failed().SetMsg("请提交用户账号或密码！").Regin(c)
			return
		}
		username := param["username"].(string)
		password := param["password"].(string)
		data, err := gf.Model("business_account").Fields("id,account_id,business_id,password,salt,name,status").Where("username", username).WhereOr("email", username).Find()
		if err != nil {
			if strings.Contains(err.Error(), " dial tcp") {
				gf.Failed().SetMsg("数据库连接失败，请检查一下数据库是否正常！").Regin(c)
			} else {
				gf.Failed().SetMsg(err.Error()).Regin(c)
			}
			return
		}
		if data == nil {
			gf.Failed().SetMsg("账号不存在！").Regin(c)
			return
		}
		if data["status"].Int() == 1 {
			gf.Failed().SetMsg("账号被禁用了").Regin(c)
			return
		}
		pass := gf.Md5(password + data["salt"].String())
		if pass != data["password"].String() {
			gf.AddloginLog(c, gf.Map{"uid": data["id"], "account_id": data["account_id"], "business_id": data["business_id"], "type": "business", "status": 1, "des": "账号登录", "error_msg": "输入的密码不正确！"})
			gf.Failed().SetMsg("您输入的密码不正确！").SetData(pass).Regin(c)
			return
		}
		loginCaptcha, _ := gcfg.Instance().Get(c, "app.loginCaptcha")
		if gf.Bool(loginCaptcha) && !gf.VerifyCaptcha(gf.String(param["codeid"]), gf.String(param["captcha"])) {
			gf.AddloginLog(c, gf.Map{"uid": data["id"], "account_id": data["account_id"], "business_id": data["business_id"], "type": "business", "status": 1, "des": "账号登录", "error_msg": "输入的验证码不正确！"})
			gf.Failed().SetMsg("您输入的验证码不正确！").Regin(c)
			return
		}
		//创建token
		token, err := gf.CreateToken(gf.Map{
			"ID":          data["id"].Int64(),
			"account_id":  data["account_id"].Int64(),
			"business_id": data["business_id"].Int64(),
		})
		if err != nil {
			gf.Failed().SetMsg(err.Error()).Regin(c)
		} else {
			gf.Model("business_account").Where("id", data["id"]).Data(map[string]interface{}{"loginstatus": 1, "last_login_time": gtime.Timestamp(), "last_login_ip": gf.GetIp(c)}).Update()
			gf.AddloginLog(c, gf.Map{"uid": data["id"], "account_id": data["account_id"], "business_id": data["business_id"], "type": "business", "status": 0, "des": "账号登录"})
			gf.Success().SetMsg("登录成功！").SetData(token).Regin(c)
		}
	} else if email, ok := param["email"]; ok {
		data, err := gf.Model("business_account").Fields("id,account_id,business_id,password,salt,name,status").Where("email", email).Find()
		if data == nil || err != nil {
			gf.Failed().SetMsg("邮箱账号不存在！").Regin(c)
			return
		}
		if data["status"].Int() == 1 {
			gf.AddloginLog(c, gf.Map{"uid": data["id"], "account_id": data["account_id"], "business_id": data["business_id"], "type": "business", "status": 1, "des": "邮箱登录", "error_msg": "账号被禁用"})
			gf.Failed().SetMsg("账号被禁用了").Regin(c)
			return
		}
		code, emerr := gf.GetVerifyCode(gf.String(param["email"]))
		if emerr != nil || code != gf.Int(param["captcha"]) {
			gf.AddloginLog(c, gf.Map{"uid": data["id"], "account_id": data["account_id"], "business_id": data["business_id"], "type": "business", "status": 1, "des": "邮箱登录", "error_msg": "验证码无效"})
			gf.Failed().SetMsg("验证码无效").SetData(emerr).Regin(c)
			return
		}
		//创建token
		token, err := gf.CreateToken(gf.Map{
			"ID":          data["id"].Int64(),
			"account_id":  data["account_id"].Int64(),
			"business_id": data["business_id"].Int64(),
		})
		if err != nil {
			gf.Failed().SetMsg(err.Error()).Regin(c)
		} else {
			gf.Model("business_account").Where("id", data["id"]).Data(map[string]interface{}{"loginstatus": 1, "last_login_time": gtime.Timestamp(), "last_login_ip": gf.GetIp(c)}).Update()
			gf.AddloginLog(c, gf.Map{"uid": data["id"], "account_id": data["account_id"], "business_id": data["business_id"], "type": "business", "status": 0, "des": "邮箱登录"})
			gf.Success().SetMsg("登录成功返回token！").SetData(token).Regin(c)
		}
	} else if mobile, ok := param["mobile"]; ok {
		data, err := gf.Model("business_account").Fields("id,account_id,business_id,password,salt,name,status").Where("mobile", mobile).Find()
		if data == nil || err != nil {
			gf.Failed().SetMsg("手机账号不存在！").Regin(c)
			return
		}
		if data["status"].Int() == 1 {
			gf.AddloginLog(c, gf.Map{"uid": data["id"], "account_id": data["account_id"], "business_id": data["business_id"], "type": "business", "status": 1, "des": "手机号登录", "error_msg": "账号被禁用"})
			gf.Failed().SetMsg("账号被禁用了").Regin(c)
			return
		}
		code, emerr := gf.GetVerifyCode(gf.String(param["mobile"]))
		if emerr != nil || code != gf.Int(param["captcha"]) {
			gf.AddloginLog(c, gf.Map{"uid": data["id"], "account_id": data["account_id"], "business_id": data["business_id"], "type": "business", "status": 1, "des": "手机号登录", "error_msg": "验证码无效"})
			gf.Failed().SetMsg("验证码无效").SetData(emerr).Regin(c)
			return
		}
		//创建token
		token, err := gf.CreateToken(gf.Map{
			"ID":          data["id"].Int64(),
			"account_id":  data["account_id"].Int64(),
			"business_id": data["business_id"].Int64(),
		})
		if err != nil {
			gf.Failed().SetMsg(err.Error()).Regin(c)
		} else {
			gf.Model("business_account").Where("id", data["id"]).Data(gf.Map{"last_login_time": gtime.Timestamp(), "last_login_ip": gf.GetIp(c)}).Update()
			gf.AddloginLog(c, gf.Map{"uid": data["id"], "account_id": data["account_id"], "business_id": data["business_id"], "type": "business", "status": 0, "des": "手机号登录"})
			gf.Success().SetMsg("登录成功！").SetData(token).Regin(c)
		}
	} else {
		gf.Failed().SetMsg("该登录方式为开发请使用其他方式登录！").Regin(c)
	}
}

/**
* 2.《获取用户》
 */
func (api *Index) GetUserinfo(c *gf.GinCtx) {
	userID := c.GetInt64("userID")
	userdata, err := gf.Model("business_account").Fields("id,business_id,name,nickname,mobile,email,avatar,status,createtime,pwd_reset_time").Where("id", userID).Find()
	if err != nil {
		gf.Failed().SetMsg("查找用户数据！").Regin(c)
	} else {
		if userdata["avatar"] == nil {
			userdata["avatar"] = gvar.New(gf.GetMainURLLocal() + "resource/uploads/static/avatar.png")
		} else if !strings.Contains(userdata["avatar"].String(), "http") {
			userdata["avatar"] = gvar.New(gf.GetFullPath(userdata["avatar"].String()))
		}
		userdata["rooturl"] = gf.VarNew(gf.GetMainURL())
		userdata["localurl"] = gf.VarNew(gf.GetMainURLLocal())
		userdata["mobile"] = gf.VarNew(gf.HideStrInfo("mobile", userdata["mobile"].String()))
		userdata["email"] = gf.VarNew(gf.HideStrInfo("email", userdata["email"].String()))
		gf.Success().SetMsg("获取用户信息").SetData(userdata).Regin(c)
	}
}

/**
*  3退出登录
 */
func (api *Index) Logout(c *gf.GinCtx) {
	user, err := gf.ParseTokenNoValid(c) //当前用户
	if err == nil {
		gf.Model("business_account").Where("id", user.ID).Data(gf.Map{"loginstatus": 0}).Update()
	}
	gf.Success().SetMsg("退出登录").SetData(true).Regin(c)
}
