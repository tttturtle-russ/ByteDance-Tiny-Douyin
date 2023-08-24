package main

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/routers"
	"log"
)

func main() {
	db.InitDb()

	router := routers.InitRouter()
	log.Fatal(router.Run())
}
