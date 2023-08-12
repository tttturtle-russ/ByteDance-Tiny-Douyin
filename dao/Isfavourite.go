package dao

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/models"
	"gorm.io/gorm"
)

func IsfavouriteByid(videoid int64) *gorm.DB {
	sql := db.DB
	var video models.Video
	find := sql.Model(&models.Video{}).Where("id= ?", videoid).First(&video)
	return find
}
