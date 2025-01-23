package routeuse

import (
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
)

func Recover(c *gin.Context) {
	defer func() {
		//记录访问日志
		if r := recover(); r != nil {
			debug.PrintStack()
			//使用Contains()函数
			erromsg := errorToString(r)
			res1 := strings.Contains(erromsg, "token")
			if res1 {
				c.JSON(http.StatusOK, gin.H{
					"code":    401, //重新登录
					"message": erromsg,
					"data":    nil,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":    1, //错误
					"message": erromsg,
					"data":    nil,
				})
			}
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
