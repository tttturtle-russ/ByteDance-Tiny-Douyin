package db

import "ByteDance-Tiny-Douyin/model"

// InitDB 用于初始化一系列数据库，Mysql,redis等
// 同时在数据库中建表，用AutoMigrate方法
func InitDB() {
	//initRedis()
	initMySql()
	models := make([]interface{}, 0)
	models = append(models, &model.User{}, &model.Video{})
	if err := MySqlDB.AutoMigrate(models...); err != nil {
		panic(err)
	}
}
