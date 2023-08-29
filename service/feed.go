package service

import (
	"ByteDance-Tiny-Douyin/model"
	"time"
)

// GetVideoByTime 调用dao层同名函数
func (svc *Service) GetVideoByTime(latestTime time.Time) (videos []model.Video, err error) {
	return svc.d.GetVideoByTime(latestTime)
}

// GetInfoByUserID 调用dao层同名函数
func (svc *Service) GetInfoByUserID(userID int64) (user model.User, err error) {
	return svc.d.GetInfoByUserID(userID)
}
