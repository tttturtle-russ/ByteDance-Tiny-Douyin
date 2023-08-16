package main

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/routers"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.SetConfigFile("./config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	db.InitDB()
	r := routers.InitRouter()
	log.Fatal(r.Run())
}
