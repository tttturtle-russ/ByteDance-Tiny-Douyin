package main

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/routers"
	"log"
)

func init() {
	//连接数据库
	db.InitDB()
}

func main() {
	router := routers.InitRouter()
	log.Fatal(router.Run())
}
