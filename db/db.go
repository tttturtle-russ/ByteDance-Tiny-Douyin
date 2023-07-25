package db

import "ByteDance-Tiny-Douyin/model"

func InitDB() {
	initRedis()
	initMySql()
	models := make([]interface{}, 0)
	models = append(models, &model.User{})
	if err := db.AutoMigrate(models...); err != nil {
		panic(err)
	}
}
