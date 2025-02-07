package controller

/**
* app路由引入口《引入模块控制器》
*
* 请把您使用包用 _ "gofly/app/xx"导入您编写的包 自动生成路由，如需要路由钩子直接引入如"gofly/app/xx"，不需路由钩子再加 _
* 不需要使用的模块则注释掉 例如home模块暂时用不到就注释掉，这样不占用资源，使用是取消注释即可。
* 路由规则：包路径“business/article” + 包中结构体“Cate”转小写+方法名（首字母转小写）
* import引入模块说明：需要添加路由钩子的直接引如gofly/app/business，如果不需路由钩子则前面加_ 如 _ "gofly/app/common"
* 有控制的模块请在RouterHandler添加模块的路由钩子
 */
import (
	"gofly/app/business"
	_ "gofly/app/common"
	"gofly/utils/gf"
)

// 路由中间件/路由钩子-需要拦截的模块再添加，如admin和business需要权限则添加，common不需要我们就加
func RouterHandler(c *gf.GinCtx) {
	business.RouterHandler(c, "business")
}
