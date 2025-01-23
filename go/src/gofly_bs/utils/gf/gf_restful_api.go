package gf

import (
	"encoding/json"
	"errors"
	"fmt"
	"gofly/utils/router/routeuse"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gtime"
	"io"
	"net/http"
	"strings"
)

const defaultMultipartMemory = 32 << 20 // 32 MB
// 组装Restful API 接口返回数据
// 返回信息主体
type (
	R struct {
		Code    int
		Message string
		Data    interface{}
		Exdata  interface{}
		Token   string
		Time    int64
	}
	UserObj = *routeuse.UserClaims
)

// 错误码
var (
	succCode = 0 // 成功
	errCode  = 1 // 失败
)

// 设置返回内容
func (r *R) SetData(data interface{}) *R {
	r.Data = data
	return r
}

// 设置返回扩展内容内容
func (r *R) SetExdata(exdata interface{}) *R {
	r.Exdata = exdata
	return r
}

// 设置编码
func (r *R) SetCode(code int) *R {
	r.Code = code
	return r
}

// 设置返回提示信息
func (r *R) SetMsg(msg string) *R {
	r.Message = msg
	return r
}

// 设置返回Token
func (r *R) SetToken(token string) *R {
	r.Token = token
	return r
}

// 返回成功内容
func Success() *R {
	r := &R{}
	r.Message = "Success"
	r.Code = succCode
	r.Time = gtime.Now().UnixMilli()
	return r
}

// 接口返回成功内容
func (r *R) Regin(ctx *GinCtx) {
	if r.Token == "" {
		//如果已经登录则刷新token
		getuser, exists := ctx.Get("user") //当前用户
		if exists {
			userinfo := getuser.(*routeuse.UserClaims)
			tokenouttime := gconv.Int64(appConf_arr["tokenouttime"])
			if userinfo.ExpiresAt-gtime.Now().Unix() < tokenouttime*60/2 { //小设置的时间超时时间一半就刷新/单位秒
				token := ctx.Request.Header.Get("Authorization")
				tockenarr, err := routeuse.Refresh(token)
				if err == nil {
					r.Token = gconv.String(tockenarr)
				}
			}
		}
	}
	ctx.JSON(http.StatusOK, GinObj{
		"code":    r.Code,
		"message": r.Message,
		"data":    r.Data,
		"exdata":  r.Exdata,
		"token":   r.Token,
		"time":    r.Time,
	})
}

// 返回失败内容
func Failed() *R {
	r := &R{}
	r.Message = "Fail"
	r.Code = errCode
	r.Data = false
	r.Time = gtime.Now().UnixMilli()
	return r
}

// 请求相关
// 批量获取请求参数-通用
func RequestParam(c *GinCtx) (map[string]interface{}, error) {
	c.Request.ParseForm()
	dataMap := make(map[string]interface{})
	if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
		if strings.Contains(c.Request.Header.Get("Content-Type"), "application/json") {
			body, _ := io.ReadAll(c.Request.Body)
			var parameter map[string]interface{}
			err := json.Unmarshal(body, &parameter)
			if err != nil {
				return nil, err
			}
			dataMap = parameter
		} else {
			//说明:须post方法,加: 'Content-Type': 'application/x-www-form-urlencoded'
			for key, valueArray := range c.Request.PostForm {
				if len(valueArray) > 1 {
					errMsg := fmt.Sprintf("#ERROR#[%s]参数设置了[%d]次,只能设置一次.", key, len(valueArray))
					return nil, errors.New(errMsg)
				}
				dataMap[key] = c.PostForm(key)
			}
		}
		//对form-data参数处理
		req := c.Request
		if err := req.ParseMultipartForm(defaultMultipartMemory); err != nil {
			if !errors.Is(err, http.ErrNotMultipart) {
				return nil, errors.New(fmt.Sprintf("error on parse multipart form array: %v", err))
			}
		}
		for key, val := range req.PostForm {
			if len(val) > 0 {
				dataMap[key] = val[0]
			}
		}
	}
	for key, _ := range c.Request.URL.Query() {
		dataMap[key] = c.Query(key)
	}
	return dataMap, nil
}

// 获取post传过来的data
func PostParam(c *GinCtx) (map[string]interface{}, error) {
	body, _ := io.ReadAll(c.Request.Body)
	var parameter map[string]interface{}
	err := json.Unmarshal(body, &parameter)
	if err != nil {
		return nil, err
	}
	return parameter, nil
}

/***Token处理****/

// 获取token用户信息
func GetUserInfo(c *GinCtx) (value *routeuse.UserClaims, exists bool) {
	getuser, exists := c.Get("user") //当前用户
	if exists {
		value = getuser.(*routeuse.UserClaims)
	} else {
		value = nil
		jwtmsg, existsmsg := c.Get("jwtmsg") //提示消息
		if existsmsg {
			fmt.Println("token失效提示", jwtmsg)
		}
	}
	return
}

// 获取token用户信息返回错误信息
func GetUserInfoErr(c *GinCtx) (value *routeuse.UserClaims, err error) {
	getuser, exists := c.Get("user") //当前用户
	if exists {
		value = getuser.(*routeuse.UserClaims)
		err = nil
	} else {
		value = nil
		jwtmsg, existsmsg := c.Get("jwtmsg") //提示消息
		if existsmsg {
			err = errors.New(jwtmsg.(string))
		} else {
			err = errors.New("登录失效，请重新登录！")
		}
	}
	return
}

// 解下Token获得用户信息
func ParseTokenGetInfo(c *GinCtx) (userinfo *routeuse.UserClaims, err error) {
	token := c.Request.Header.Get("Authorization")
	userinfo, err = routeuse.ParseToken(token)
	return
}

// 创建登录Token
func CreateToken(param Map) (token string, err error) {
	var Mid int64
	var accountID int64
	var businessID int64
	if GID, ok := param["ID"]; ok {
		Mid = Int64(GID)
	}
	if MID, ok := param["id"]; ok {
		Mid = Int64(MID)
	}
	if accountIDv, ok := param["account_id"]; ok {
		accountID = Int64(accountIDv)
	}
	if businessIDv, ok := param["business_id"]; ok {
		businessID = Int64(businessIDv)
	}
	token, err = routeuse.GenerateToken(&routeuse.UserClaims{ID: Mid, AccountID: accountID, BusinessID: businessID})
	return
}

// 解下Token获得用户信息拦截错误提示
func ParseTokenMsg(c *GinCtx) *routeuse.UserClaims {
	token := c.Request.Header.Get("Authorization")
	userinfo, err := routeuse.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, GinObj{
			"code":    401,
			"message": err,
			"data":    "",
			"exdata":  "token验证失败",
			"token":   "",
			"time":    gtime.Now().UnixMilli(),
		})
		c.Abort()
		return nil
	} else {
		return userinfo
	}
}

// 仅获取token信息-不需验证
func ParseTokenNoValid(c *GinCtx) (userinfo *routeuse.UserClaims, err error) {
	token := c.Request.Header.Get("Authorization")
	userinfo, err = routeuse.ParseJwt(token)
	return
}
