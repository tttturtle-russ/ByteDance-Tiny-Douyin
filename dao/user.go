package dao

import (
	"ByteDance-Tiny-Douyin/model"
)

func (d *Dao) RegisterUser(user model.User) (uint, error) {
	err := d.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return 0, err
	}
	//err = d.Model(user).Where(&user).Select("id").Scan(&id).Error
	return user.ID, err
}
