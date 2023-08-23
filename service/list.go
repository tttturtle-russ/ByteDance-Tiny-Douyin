package service

import "ByteDance-Tiny-Douyin/model"

// GetVideosByID 调用dao层同名函数
func (svc *Service) GetVideosByID(userID string) (videos []model.Video, err error) {
	return svc.d.GetVideosByID(userID)
}
