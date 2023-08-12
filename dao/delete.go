package dao

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/models"
	"gorm.io/gorm"
)

// 删除video列表中的user
func Delete(userid int64, videoid int64) error {
	err := db.DB.Model(&models.Video{}).Where("video_id = ?", videoid).Update("user_favourites", gorm.Expr("array_remove(user_favourites, ?)", userid)).Error
	return err
}

// 减少video总点赞数
func TotalDown(videoid int64) error {
	err := db.DB.Model(&models.Video{}).Where("id = ?", videoid).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	return err
}

// 减少video作者的获赞数
func AuthorDown(videoid int64) error {
	var video models.Video
	err := db.DB.Model(&models.Video{}).Where("id = ?", videoid).First(&video).Error
	if err != nil {
		return err
	}

	authorid := video.Author.Id
	err = db.DB.Model(&models.User{}).Where("id = ?", authorid).UpdateColumn("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error
	return err
}

// 减少用户的点赞数
func UserDown(userid int64) error {
	err := db.DB.Model(&models.User{}).Where("id = ?", userid).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	return err
}
