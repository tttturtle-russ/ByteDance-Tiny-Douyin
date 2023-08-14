package main

import (
	"GoWorkspace/db"
	"GoWorkspace/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//获取初始化的数据库
	realdb := db.InitDB()
	//延迟关闭数据库
	defer realdb.Close()
	//创建一个默认的路由引擎
	r := gin.Default()
	//启动路由
	routes.CollectRoutes(r)

	//在9090端口启动服务
	panic(r.Run(":9090"))
}
