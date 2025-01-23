package business

/**
* 引入控制器-文件夹名称的路径
* 请把您使用包用 _ "gofly/app/business/XX"导入您编写的包 自动生成路由
* 不是使用则注释掉
* 即：http://xx.com/business/article/cate/getList
* 如果文件夹内没有对应package的go文件请把控制器类删除
 */
import (
	_ "gofly/app/business/common"
	_ "gofly/app/business/dashboard"
	_ "gofly/app/business/datacenter"
	_ "gofly/app/business/developer"
	_ "gofly/app/business/system"
	_ "gofly/app/business/user"
	"gofly/utils/gf"
	_ "gofly/app/business/createcode"
)

// 路由中间件/路由钩子，noAuths无需路由验证接口，可以从c获取请求各种参数
func RouterHandler(c *gf.GinCtx, modelname string) {
	if gf.IsModelPath(c.FullPath(), modelname) { //在这里面处理拦截操作，如果需要拦截终止执行则：c.Abort()
		// 判断请求接口是否需要验证权限(RBAC的权限)
		if gf.NeedAuthMatch(c) {
			haseauth := gf.CheckAuth(c, modelname)
			if haseauth {
				c.Next()
			} else {
				gf.Failed().SetMsg(gf.LocaleMsg().SetLanguage(c.Request.Header.Get("locale")).Message("sys_auth_permission")).SetData(haseauth).Regin(c)
				c.Abort()
			}
		} else {
			c.Next()
		}
	}
}
