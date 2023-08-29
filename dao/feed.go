package dao

import (
	"ByteDance-Tiny-Douyin/model"
	"time"
)

// GetVideoByTime 根据视频发布时间获取视频信息
func (d *Dao) GetVideoByTime(latestTime time.Time) (videos []model.Video, err error) {
	const videoMaxNum int = 10
	err = d.Where("created_at < ?", latestTime).Order("created_at desc").Limit(videoMaxNum).Find(&videos).Error
	return videos, err
}

// GetInfoByUserID 根据用户ID获取用户信息
func (d *Dao) GetInfoByUserID(userID int64) (user model.User, err error) {
	err = d.Where("id = ?", userID).First(&user).Error
	return user, err
}
