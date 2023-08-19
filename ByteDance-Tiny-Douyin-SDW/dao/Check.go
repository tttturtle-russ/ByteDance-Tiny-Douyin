package dao

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/models"
)

func CheckFollow(userid int64, followid int64) (bool, error) {
	var follow models.FollowList
	err := db.DB.Model(&models.FollowList{}).Where("user_id = ? AND follow_id = ?", userid, followid).First(&follow)

	if err != nil {
		return false, err.Error //查询记录不存在或查询出错
	} else {
		return true, err.Error
	}
}
