package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	Id         int64     `json:"id"`
	ToUserId   int64     `json:"to_user_id" binding:"required"`
	FromUserId int64     `json:"from_user_id" binding:"required"`
	Content    string    `json:"content" binding:"required"`
	CreateAt   time.Time `json:"create_time,string"`
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	m.CreateAt = time.Now()
	return nil
}
