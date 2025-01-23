package gf

/**
* 自动路由工具
 */
import (
	"gofly/utils/tools/gstr"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

// 路由结构体
type Route struct {
	Path        string         //url路径
	HttpMethod  string         //http方法 get post
	Method      reflect.Value  //方法路由
	Args        []reflect.Type //参数类型
	Action      string         //函数名
	NoNeedLogin []string       //无需登录的方法，全部忽略则为["*"]
	NoNeedAuths []string       //无需接口权限方法，全部忽略则为["*"]
}

// 路由集合
var Routes = []Route{}

// 注册控制器
func Register(controller, fpath interface{}) bool {
	PkgPathstr := reflect.TypeOf(fpath).PkgPath() //获取文件路径
	vbf := reflect.ValueOf(controller)
	var Ruledata = vbf.Elem()
	var noNeedLoginSlice []string = make([]string, 0) //无需要验证权限
	if Ruledata.NumField() > 0 && Ruledata.FieldByName("NoNeedLogin").IsValid() {
		noNeedLoginSlice = Strings(Ruledata.FieldByName("NoNeedLogin"))
	}
	var noNeedAuthsSlice []string = make([]string, 0) //无需要验证权限
	if Ruledata.NumField() > 0 && Ruledata.FieldByName("NoNeedAuths").IsValid() {
		noNeedAuthsSlice = Strings(Ruledata.FieldByName("NoNeedAuths"))
	}
	//非控制器或无方法则直接返回
	if vbf.NumMethod() == 0 {
		return false
	}
	rootPkg := ""
	if strings.Contains(PkgPathstr, "/app/") {
		PkgPath_arr := strings.Split(PkgPathstr, "/app/")
		rootPkg = PkgPath_arr[len(PkgPath_arr)-1]
	}
	ctrlName := reflect.TypeOf(controller).String()
	module := ctrlName
	if strings.Contains(ctrlName, ".") {
		module = ctrlName[strings.Index(ctrlName, ".")+1:]
	}
	if module == "Index" { //去index
		module = "/"
	} else {
		module = "/" + strings.ToLower(module) + "/"
	}
	//遍历方法
	for i := 0; i < vbf.NumMethod(); i++ {
		method := vbf.Method(i)
		action := vbf.Type().Method(i).Name
		//拼接路由地址
		path := rootPkg + module + gstr.LcFirst(action)
		//遍历参数
		params := make([]reflect.Type, 0, vbf.NumMethod())
		httpMethod := "POST" //默认POST
		if (strings.HasPrefix(action, "Get") && !strings.HasPrefix(action, "GetPost")) || action == "Index" {
			httpMethod = "GET"
		} else if strings.HasPrefix(action, "Del") || action == "Del" {
			httpMethod = "DELETE"
		} else if strings.HasPrefix(action, "Put") || action == "Put" {
			httpMethod = "PUT"
		}
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
		}
		if strings.HasSuffix(action, "ParId") {
			path = path + "/:id"
		}
		route := Route{Path: path, Method: method, Args: params, Action: action, NoNeedLogin: noNeedLoginSlice, NoNeedAuths: noNeedAuthsSlice, HttpMethod: httpMethod}
		Routes = append(Routes, route)
		if strings.HasPrefix(action, "GetPost") { //再增加一个get请求
			route := Route{Path: path, Method: method, Args: params, Action: action, NoNeedLogin: noNeedLoginSlice, NoNeedAuths: noNeedAuthsSlice, HttpMethod: "GET"}
			Routes = append(Routes, route)
		}
	}
	return true
}

// 根据path匹配对应的方法
func PathMatch(path string, route Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		fields := strings.Split(path, "/")
		if len(fields) < 2 { //路径最少两层，少于两层将无法注册
			return
		}
		if len(Routes) > 0 {
			arguments := make([]reflect.Value, 1)
			arguments[0] = reflect.ValueOf(c) // *gin.Context
			route.Method.Call(arguments)
		}
	}
}

// 检测当前控制器和方法是否需要权限认证-需要=true,不需要=false
// Action是请求方法，NoNeedAuths需要验证权限的数组
func NeedAuthMatch(c *GinCtx) bool {
	Actionstr := gstr.LcFirst(c.GetString("Action"))
	for _, eachItem := range c.GetStringSlice("NoNeedAuths") {
		if gstr.LcFirst(eachItem) == Actionstr || eachItem == "*" {
			return false
		}
	}
	return true
}

// 判断是否需要登录验证-需要=true,不需要=false
func NeedLoginMatch(actiond string, noNeedLoginMap []string) bool {
	Actionstr := gstr.LcFirst(actiond)
	for _, eachItem := range noNeedLoginMap {
		if gstr.LcFirst(eachItem) == Actionstr || eachItem == "*" {
			return false
		}
	}
	return true
}
