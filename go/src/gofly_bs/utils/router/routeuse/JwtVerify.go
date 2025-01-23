package routeuse

import (
	"errors"
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gctx"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 用户信息类，作为生成token的参数
type UserClaims struct {
	ID         int64 `json:"id"`          //用户数据id
	AccountID  int64 `json:"account_id"`  //A、B端用户账号id
	BusinessID int64 `json:"business_id"` //B端主账号id
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	ctx         = gctx.New()
	appConf, _  = gcfg.Instance().Get(ctx, "app")
	appConf_arr = gconv.Map(appConf)
	//自定义的token秘钥 tokensecret
	secret     = []byte(gconv.String(appConf_arr["tokensecret"]))
	effectTime = time.Duration(gconv.Int64(appConf_arr["tokenouttime"])) * time.Minute //单位分钟
)

// 生成token
func GenerateToken(claims *UserClaims) (sign string, err error) {
	//设置token有效期，也可不设置有效期，采用redis的方式
	//   1)将token存储在redis中，设置过期时间，token如没过期，则自动刷新redis过期时间，
	//   2)通过这种方式，可以很方便的为token续期，而且也可以实现长时间不登录的话，强制登录
	//本例只是简单采用 设置token有效期的方式，只是提供了刷新token的方法，并没有做续期处理的逻辑
	claims.StandardClaims = jwt.StandardClaims{} //jwt-go提供的标准claim
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	//生成token
	sign, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	return
}

// 验证token
func JwtVerify(c *gin.Context) {
	//获取token字符串
	token := c.GetHeader("Authorization")
	if token == "" {
		token = c.GetHeader("authorization")
	}
	if token != "" {
		//验证token，并存储在请求中
		tokeninfo, err := ParseToken(token)
		if err != nil {
			c.Set("jwtempty", true)
			c.Set("jwtmsg", err.Error())
		} else {
			c.Set("jwtempty", false)
			c.Set("user", tokeninfo)
			c.Set("userID", tokeninfo.ID)
			c.Set("accountID", tokeninfo.AccountID)
			c.Set("businessID", tokeninfo.BusinessID)
		}
	} else {
		c.Set("jwtempty", true)
		c.Set("jwtmsg", "token 不存在")
	}
}

// 解析Token-验证过去
func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("The token is invalid")
	}
	return claims, err
}

// 仅仅解析数据-不验证
func ParseJwt(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if token == nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("The token is invalid")
	}
	return claims, nil
}

// 更新token
func Refresh(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return "", errors.New("The token is invalid")
	}
	jwt.TimeFunc = time.Now
	claims.StandardClaims.ExpiresAt = time.Now().Add(effectTime).Unix()
	sign, serr := GenerateToken(claims) //生成token
	return sign, serr
}
