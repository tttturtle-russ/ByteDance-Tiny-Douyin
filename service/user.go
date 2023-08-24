package service

import (
	"ByteDance-Tiny-Douyin/model"
)

func (svc *Service) RegisterUser(username, password string) (uint, error) {
	user := model.User{
		Name:     username,
		Password: password,
	}
	return svc.d.RegisterUser(user)
}
func (svc *Service) LoginUser(username, password string) (uint, error) {
	user := model.User{
		Name:     username,
		Password: password,
	}
	return svc.d.LoginUser(user)
}
