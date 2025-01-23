package main

import (
	//引入数据库驱动-去https://doc.goflys.cn/docview?id=26&fid=395下载
	_ "gofly/utils/drivers/pgsql"
	// _ "gofly/utils/drivers/redis"
	"gofly/utils/router"
)

func main() {
	// 启动服务器
	router.RunServer()
}
