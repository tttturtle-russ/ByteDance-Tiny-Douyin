package service

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/model"
)

func RegisterUser(username, password string) error {
	user := model.User{
		Name:     username,
		Password: password,
	}
	return dao.RegisterUser(user)
}
