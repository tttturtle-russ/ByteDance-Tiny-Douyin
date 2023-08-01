package service

import "ByteDance-Tiny-Douyin/model"

func (svc *Service) GetMessageList(fromUserId, toUserId int64) (list []model.Message, err error) {
	return svc.d.GetMessageList(fromUserId, toUserId)
}
