package main

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/routers"
	"log"
)

func main() {
	//连接数据库
	err := dao.InitMysql()
	if err != nil {
		panic(err.Error())
	}
	log.Println("database connect succeed")

	router := routers.InitRouter()
	log.Fatal(router.Run())
}
