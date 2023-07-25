package dao

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/model"
)

func RegisterUser(user model.User) error {
	return db.Create(&user).Error
}
