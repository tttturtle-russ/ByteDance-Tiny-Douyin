package dao

import (
	"ByteDance-Tiny-Douyin/model"
	"gorm.io/gorm"
)

type Dao struct {
	*gorm.DB
}

// NewDao 创建一个Dao实例
func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db.Session(&gorm.Session{}),
	}
}

// GetVideosByID 根据传入的用户ID获取该用户发布的所有视频信息
func GetVideosByID(DB *Dao, userID string, videos *[]model.Video) (err error) {
	err = DB.Where("author_id = ?", userID).Find(videos).Error
	return err
}

// GetUserInfoByID 根据传入的用户ID获取该用户信息
func GetUserInfoByID(DB *Dao, userID string, user *model.User) (err error) {
	err = DB.Where("id = ?", userID).First(&user).Error
	return err
}
