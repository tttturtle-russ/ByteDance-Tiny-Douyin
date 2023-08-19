package Db

import (
	"ByteDance-Tiny-Douyin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() error {
	var err error
	dsn := "root:1234@tcp(127.0.0.1:13306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	DB.AutoMigrate(&models.User{}, &models.Video{})
	return nil
}
