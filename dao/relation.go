package dao

import "ByteDance-Tiny-Douyin/model"

func (d *Dao) GetFriendList(id uint) ([]model.User, error) {
	var list []model.User
	err := d.Model(&model.User{}).
		Where().Error
}
