package db

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySqlDB *gorm.DB

// initMySql 先从config.json 中读取数据库信息，然后建立连接
func initMySql() {
	var err error
	viper.SetConfigFile("E:\\ByteDance_Tiktok\\ByteDance-Tiny-Douyin\\config.json")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed to read config file: %s", err))
	}
	username := viper.GetString("database.mysql.username")
	password := viper.GetString("database.mysql.password")
	host := viper.GetString("database.mysql.host")
	port := viper.GetInt("database.mysql.port")
	dbname := viper.GetString("database.mysql.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	MySqlDB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
}
