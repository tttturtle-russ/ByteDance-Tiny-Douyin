package service

import "ByteDance-Tiny-Douyin/model"

func (svc *Service) GetuserInfo(username string) (model.User, error) {
	user := model.User{
		Name: username,
	}
	return svc.d.Userinfo(user)
}
