package dao

import (
	"ByteDance-Tiny-Douyin/model"
)

func (d *Dao) Userinfo(user model.User) (model.User, error) {
	name := user.Name
	client := model.User{}
	err := d.Where("Name = ?", name).First(&client).Error
	if client.ID == 0 || err != nil {
		return model.User{}, err
	}
	return client, nil
}
