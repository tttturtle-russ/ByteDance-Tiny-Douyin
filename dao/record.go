package dao

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/models"
)

// 增加LikeList记录
func RecordAdd(userid int64, videoid int64) error {
	add := &models.LikeList{
		UserId:  userid,
		VideoId: videoid,
	}

	result := db.DB.Model(&models.LikeList{}).Create(add)
	return result.Error
}

// 删除LikeList记录
func RecordDelete(userid int64, videoid int64) error {
	result := db.DB.Model(&models.LikeList{}).Where("user_id = ? AND video_id = ?", userid, videoid).Delete(&models.LikeList{})
	return result.Error
}
