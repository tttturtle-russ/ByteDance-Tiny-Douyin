package service

import (
	"ByteDance-Tiny-Douyin/model"
	"errors"
)

func (svc *Service) GetMessageList(fromUserId uint, toUserId uint) (list []model.Message, err error) {
	return svc.d.GetMessageList(fromUserId, toUserId)
}

func (svc *Service) SendMessage(fromUser uint, toUserId uint, content string, actionType int) error {
	switch actionType {
	case 1:
		return svc.d.SendMessage(fromUser, toUserId, content)
	default:
		return errors.New("not support action type")
	}
}
