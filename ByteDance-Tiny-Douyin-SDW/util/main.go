package main

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/routers"
	"log"
)

func main() {
	err := db.InitMysql()
	if err == nil {
		log.Println("database connect succeed")
	} else {
		panic(err.Error())
	}

	router := routers.InitRouter()
	log.Fatal(router.Run())
}
