package gf

import (
	"image/color"
	"time"

	"github.com/mojocn/base64Captcha"
)

type CaptchaResult struct {
	Id          string `json:"id"`
	Show        bool   `json:"show"`
	Base64Blog  string `json:"img"`
	VerifyValue string `json:"code"`
	ExpireTime  int64  `json:"expireTime"`
}

// 默认存储10240个验证码，每个验证码3分钟过期
// Expiration time of captchas used by default store.
var Expiration = 3 * time.Minute
var store = base64Captcha.NewMemoryStore(10240, Expiration)

// 生成图片验证码
func GenerateCaptcha() (interface{}, error) {
	// 生成默认数字
	driver := base64Captcha.NewDriverMath(39, 110, 0, 0, &color.RGBA{0, 0, 0, 1}, nil, []string{"wqy-microhei.ttc"})
	// 生成base64图片
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 获取
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		return "", err
	}
	captchaResult := CaptchaResult{Id: id, Show: Bool(appConf_arr["loginCaptcha"]), Base64Blog: b64s, ExpireTime: time.Now().Add(Expiration).UnixMilli()}
	return captchaResult, nil
}

// 校验图片验证码,并清除内存空间
func VerifyCaptcha(id string, value string) bool {
	// TODO 只要id存在，就会校验并清除，无论校验的值是否成功, 所以同一id只能校验一次
	// 注意：id,b64s是空 也会返回true 需要在加判断
	verifyResult := store.Verify(id, value, true)
	return verifyResult
}
