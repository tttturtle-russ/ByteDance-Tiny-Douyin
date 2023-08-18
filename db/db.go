package db

import "ByteDance-Tiny-Douyin/model"

func InitDb() {
	//InitMysql完成连接数据库，完成自动迁移
	InitMysql()
	err := DB.AutoMigrate(&model.User{}, &model.Video{}, &model.LikeList{})
	if err != nil {
		panic(err)
	}
}
