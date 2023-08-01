package dao

import "ByteDance-Tiny-Douyin/model"

func (d *Dao) GetMessageList(fromUserId int64, toUserId int64) ([]model.Message, error) {
	var list []model.Message
	err := d.Model(&model.Message{}).
		Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).
		Order("created_at desc").
		Find(&list).Error
	return list, err
}
