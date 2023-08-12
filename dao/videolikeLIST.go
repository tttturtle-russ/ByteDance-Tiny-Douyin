package dao

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/models"
)

func LikeList(userid int64) ([]models.Video, error) {
	var response []models.Video
	err := db.DB.Model(&models.User{}).Preload("video_favourites").Where("id = ?", userid).Find(&response).Error
	return response, err
}
