package main

import (
	"ByteDance-Tiny-Douyin/routers"
	"log"
)

func main() {
	router := routers.InitRouter()
	log.Fatal(router.Run())
}
