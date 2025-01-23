package routeuse

import (
	"gofly/utils/tools/gcfg"
	"gofly/utils/tools/gconv"
	"gofly/utils/tools/gctx"
	"net/http"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
)

func LimitHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx = gctx.New()
		appConf, _ := gcfg.Instance().Get(ctx, "app")
		appConf_arr := gconv.Map(appConf)
		// 创建过期配置
		expiraOption := &limiter.ExpirableOptions{
			// 默认过期时间设置为 1 秒
			DefaultExpirationTTL: 1 * time.Second,
		}
		lmt := tollbooth.NewLimiter(gconv.Float64(appConf_arr["limiterMax"]), expiraOption)
		lmt.SetMessage("您访问过于频繁，系统安全检查认为恶意攻击。")
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    1,
				"message": "您的操作太频繁，请稍后再试！",
				"data":    nil,
			})
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
