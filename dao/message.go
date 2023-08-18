package dao

import (
	"ByteDance-Tiny-Douyin/model"
)

func (d *Dao) GetMessageList(fromUserId uint, toUserId uint) ([]model.Message, error) {
	var list []model.Message
	err := d.Model(&model.Message{}).
		Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).
		Order("created_at desc").
		Find(&list).Error
	return list, err
}

func (d *Dao) SendMessage(fromUserId uint, toUserId uint, content string) error {
	err := d.Model(&model.Message{}).
		Create(&model.Message{
			ToUserId:   toUserId,
			FromUserId: fromUserId,
			Content:    content,
		}).Error
	return err
}
