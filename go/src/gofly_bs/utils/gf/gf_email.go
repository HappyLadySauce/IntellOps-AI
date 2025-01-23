package gf

import (
	"errors"
	"gofly/utils/tools/grand"
	"strings"

	"gopkg.in/gomail.v2"
)

// 发送邮件
// 请求参数：email邮箱地址，title邮件标题(如果为空则默认则从配置获取)，text邮件内容(如果为空则默认则从配置获取)
// 返回参数：bool 结果, error 错误提示
func SendEmail(c *GinCtx, email []string, title, text string) (bool, error) {
	if len(email) == 0 {
		return false, errors.New("请填写邮箱")
	} else {
		businessID, isbusiness := c.Get("businessID") //当前businessID
		var emailConfig OrmRecord
		if isbusiness {
			emailConfig, _ = Model("common_email").Where("data_from", "business").Where("business_id", businessID).Find()
		}
		if emailConfig == nil { //如果商务端没有配置则用管理后台发送
			emailConfig, _ = Model("common_email").Where("data_from", "common").Find()
		}
		if emailConfig == nil {
			return false, errors.New("请到业务端后台“配置管理”配置邮箱")
		} else {
			sender := emailConfig["sender_email"].String()  //发送者邮箱
			authCode := emailConfig["auth_code"].String()   //邮箱授权码
			mailTitle := emailConfig["mail_title"].String() //邮件标题
			mailBody := emailConfig["mail_body"].String()   //邮件内容,可以是html
			if title != "" {
				mailTitle = title //邮件标题
			}
			if text == "" {
				code := grand.Digits(6)
				mailBody = strings.Replace(mailBody, "{code}", code, 1)
				for _, val := range email {
					SetVerifyCode(val, code) //验证码存在本地缓存
				}
			} else {
				mailBody = text
			}
			m := gomail.NewMessage()
			m.SetHeader("From", sender)       //发送者邮箱账号
			m.SetHeader("To", email...)       //接收者邮箱列表
			m.SetHeader("Subject", mailTitle) //邮件标题
			m.SetBody("text/html", mailBody)  //邮件内容,可以是html
			//服务器地址和端口是默认腾讯的
			service_host := "smtp.qq.com"
			if _, ok := emailConfig["service_host"]; ok {
				service_host = emailConfig["service_host"].String()
			}
			service_port := 587
			if _, ok := emailConfig["service_port"]; ok {
				service_port = Int(emailConfig["service_port"])
			}
			d := gomail.NewDialer(service_host, service_port, sender, authCode)
			err := d.DialAndSend(m)
			if err != nil {
				return false, err
			} else {
				return true, nil
			}
		}
	}
}
