package dao

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/models"
	"gorm.io/gorm"
)

// video的用户列表增加
func Insert(userid int64, videoid int64) error {
	err := db.DB.Model(&models.Video{}).Where("video_id = ?", videoid).Update("user_favourites", gorm.Expr("array_append(user_favourites, ?)", userid)).Error
	return err
}

// video对应的总获赞数
func TotalAdd(videoid int64) error {
	err := db.DB.Model(&models.Video{}).Where("id = ?", videoid).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	return err
}

// video作者的获赞数
func AuthorAdd(videoid int64) error {
	var video models.Video
	err := db.DB.Model(&models.Video{}).Where("id = ?", videoid).First(&video).Error
	if err != nil {
		return err
	}

	authorid := video.Author.Id
	err = db.DB.Model(&models.User{}).Where("id = ?", authorid).UpdateColumn("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error
	if err != nil {
		return err
	}
	return err
}

// 用户的点赞数
func UserAdd(userid int64) error {
	err := db.DB.Model(&models.User{}).Where("id = ?", userid).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	return err
}
