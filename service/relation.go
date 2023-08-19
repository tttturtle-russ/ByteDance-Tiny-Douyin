package service

import "ByteDance-Tiny-Douyin/model"

func (svc *Service) GetFriendList(id uint) ([]model.User, error) {
	return svc.d.GetFriendList(id)
}
