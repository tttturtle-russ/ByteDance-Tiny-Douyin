package service

import "ByteDance-Tiny-Douyin/model"

func (svc *Service) GetFriendList(id int64) ([]model.User, error) {
	return svc.d.GetFriendList(id)
}
