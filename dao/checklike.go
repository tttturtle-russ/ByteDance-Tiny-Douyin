package dao

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/models"
)

func CheckLike(userid int64, videoid int64) (bool, error) {
	var like models.LikeList
	err := db.DB.Model(&models.LikeList{}).Where("user_id = ? AND video_id = ?", userid, videoid).First(&like)

	if err != nil {
		return false, err.Error //查询记录不存在或查询出错
	} else {
		return true, err.Error
	}
}
