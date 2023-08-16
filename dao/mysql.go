package dao

import (
	"ByteDance-Tiny-Douyin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&model.Video{}, &model.User{})
	return err
}
