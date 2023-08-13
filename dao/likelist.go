package dao

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/models"
)

func LikeList(userid int64) ([]models.Video, error) {
	//根据like_list表查找所有的video_id
	var lists []models.LikeList
	var ids []int64
	result := db.DB.Model(&models.LikeList{}).Select("video_id").Where("user_id = ?", userid).Find(&lists)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, list := range lists {
		ids = append(ids, list.VideoId)
	}

	//由video_id返回所有的video
	var videos []models.Video
	out := db.DB.Model(&models.Video{}).Where("id IN ?", ids).Find(&videos)
	if out.Error != nil {
		return nil, out.Error
	}
	return videos, out.Error
}
