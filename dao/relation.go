package dao

import "ByteDance-Tiny-Douyin/model"

func (d *Dao) GetFriendList(id int64) ([]model.User, error) {
	var list []model.User
	err := d.Model(&model.User{}).
		Where("id1 = ?", id).
		Or("id2 = ?", id).
		Find(&list).Error
	return list, err
}
