package basetool

import (
	"gofly/utils/gf"
	"gofly/utils/tools/grand"
	"time"
)

// 公共工具
type Index struct {
	NoNeedLogin []string
	NoNeedAuths []string
}

func init() {
	fpath := Index{NoNeedLogin: []string{"getCaptcha"}, NoNeedAuths: []string{"checkHaseRule"}}
	gf.Register(&fpath, fpath)
}

/**
*  1获取验证码、图片验证码、手机验证码、邮箱验证码
 */
func (api *Index) GetCaptcha(c *gf.GinCtx) {
	typeStr := c.DefaultQuery("type", "")
	switch typeStr {
	case "image": //获取图片验证码
		data, err := gf.GenerateCaptcha()
		if err != nil {
			gf.Failed().SetMsg("获取图片验证码失败").Regin(c)
		} else {
			gf.Success().SetMsg("获取图片验证码成功").SetData(data).Regin(c)
		}
	case "mobile": //获取手机验证码
		mobile := c.DefaultQuery("mobile", "")
		if mobile == "" {
			gf.Failed().SetMsg("请填写手机").Regin(c)
		} else {
			code := grand.Digits(6)
			gf.SetVerifyCode(mobile, code) //验证码存在本地缓存中
			gf.Success().SetMsg("获取手机验证码").SetData(code).Regin(c)
		}
	case "email": //获取邮箱验证码
		email := c.DefaultQuery("email", "")
		if email == "" {
			gf.Failed().SetMsg("请填写邮箱").Regin(c)
		} else {
			res, erro := gf.SendEmail(c, []string{email}, "", "")
			if erro == nil {
				gf.Success().SetMsg("获取验证码").SetData(res).Regin(c)
			} else {
				gf.Failed().SetMsg(erro.Error()).Regin(c)
			}
		}
	case "qrcode": //扫码登录
		keyval := c.DefaultQuery("keyval", "")
		if keyval == "" {
			gf.Failed().SetMsg("选择扫码类型").Regin(c)
		} else {
			addtime, _ := time.ParseDuration("60s")
			expiretime := time.Now().Add(addtime).UnixMilli()
			gf.Success().SetMsg("获取扫码内容").SetData(gf.Map{"url": keyval, "expiretime": expiretime}).Regin(c)
		}
	}
}

// 1 获取使用字典数据接口
func (api *Index) GetDicData(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	tablename, _ := gf.Model("common_dictionary_group").Where("id", param["group_id"]).Value("tablename")
	getfield := "id,keyname as label,keyvalue as value"
	if gf.DbHaseField(tablename.String(), "tagcolor") {
		getfield = "id,keyname as label,keyvalue as value,tagcolor as color"
	}
	list, err := gf.Model(tablename.String()).Where("group_id", param["group_id"]).Where("status", 0).Fields(getfield).Select()
	if err != nil {
		gf.Failed().SetMsg("获取字典数据失败！").SetData(err).Regin(c)
	} else {
		gf.Success().SetMsg("获取字典数据列表").SetData(list).Regin(c)
	}
}

// 检查路由是否已存在-因为路由不能重复
func (api *Index) CheckHaseRule(c *gf.GinCtx) {
	param, _ := gf.RequestParam(c)
	auth_rule_talbename := "business_auth_rule"
	if param["codelocation"] == "adminDirName" {
		auth_rule_talbename = "admin_auth_rule"
	}
	res, err := gf.Model(auth_rule_talbename).Where("id !=?", param["id"]).WhereIn("routename", param["routename"]).Value("id")
	if err != nil {
		gf.Failed().SetMsg("检查路由失败").SetData(err).Regin(c)
	} else {
		hase := false
		if res == nil { //可以用
			hase = true
		}
		gf.Success().SetMsg("检查路由成功！").SetData(hase).Regin(c)
	}
}
