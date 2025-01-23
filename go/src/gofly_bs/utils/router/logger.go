package router

//保存操作日志记录到数据库
import (
	"bytes"
	"gofly/utils/gf"
	"gofly/utils/plugin"
	"gofly/utils/router/routeuse"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 保存日志记录到
// CustomResponseWriter 封装 gin ResponseWriter 用于获取回包内容。
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 忽略操作日志请求接口
var (
	noSaveUrl = []string{"/user/login", "/user/logout", "/system/log/", "/system/dept/getParen", "/system/account/getRole", "/system/rule/getRoutes", "/system/rule/getParent", "/user/getUserinfo", "/user/account/getMenu", "/dashboard/workplace/getQuick"}
)

// 日志中间件。
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := c.FullPath()
		//目前日志只监听admin、business、common模块，如果新增的模块需要记录操作日志，请求自行添加。
		if GetNoInUrls(url) && (strings.HasPrefix(url, "/admin/") || strings.HasPrefix(url, "/business/") || strings.HasPrefix(url, "/common/")) {
			// 记录请求时间
			start := time.Now()
			// 使用自定义 ResponseWriter-解决多次读取响应 Body 的问题
			crw := &CustomResponseWriter{
				body:           bytes.NewBufferString(""),
				ResponseWriter: c.Writer,
			}
			c.Writer = crw

			// 打印请求信息
			var reqBodyStr interface{}
			if c.Request.Method == "GET" {
				dataMap := make(map[string]interface{})
				for key, _ := range c.Request.URL.Query() {
					dataMap[key] = c.Query(key)
				}
				reqBodyStr = dataMap
			} else {
				reqBody, _ := c.GetRawData()
				// 请求包体写回-解决多次读取请求 Body 的问题
				if len(reqBody) > 0 {
					c.Request.Body = io.NopCloser(bytes.NewBuffer(reqBody))
				}
				reqBodyStr = string(reqBody)
			}

			// 执行请求处理程序和其他中间件函数
			c.Next()

			// 记录回包内容和处理时间
			end := time.Now()
			latency := end.Sub(start)
			//操作日志入库
			getuser, exists := c.Get("user")
			if exists { //对登录用户进行操作日志保存
				user := getuser.(*routeuse.UserClaims)
				ip := c.ClientIP()
				address, _ := plugin.NewIpRegion(ip)
				savedata := gf.Map{
					"uid":            user.ID,
					"account_id":     user.AccountID,
					"business_id":    user.BusinessID,
					"request_method": c.Request.Method,
					"url":            url,
					"ip":             ip,
					"address":        address,
					"status":         crw.Status(),
					"req_headers":    c.Request.Header,
					"req_body":       reqBodyStr,
					"resp_body":      crw.body.Bytes(),
					"resp_headers":   crw.Header(),
					"latency":        latency.String(),
				}
				url_arr := strings.Split(url, "/")
				if len(url_arr) > 1 {
					savedata["type"] = url_arr[1]
					prefixTableName := "admin"
					businessID := c.GetInt64("businessID") //当前商户ID
					if businessID > 0 {
						prefixTableName = "business"
					}
					authdata, err := gf.Model(prefixTableName+"_auth_rule").Where("path", c.FullPath()).Fields("pid,title,des").Find()
					if err == nil && !authdata.IsEmpty() {
						ftitle, _ := gf.Model(prefixTableName+"_auth_rule").Where("id", authdata["pid"]).Value("title")
						substr := authdata["des"].String()
						if authdata["des"].IsEmpty() {
							substr = authdata["title"].String()
						}
						savedata["des"] = ftitle.String() + "【" + substr + "】"
					}
				}
				gf.Model("common_sys_operation_log").Insert(savedata)
			}

		} else {
			c.Next()
		}
	}
}

// 过滤忽略请求
func GetNoInUrls(url string) bool {
	for _, val := range noSaveUrl {
		if strings.Contains(url, val) {
			return false
		}
	}
	return true
}
