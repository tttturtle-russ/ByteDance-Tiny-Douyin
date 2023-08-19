package dao

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/models"

	"gorm.io/gorm"
)

// 粉丝数增加
func FollowerAdd(followid int64) error {
	err := db.DB.Model(&models.User{}).Where("id = ?", followid).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error
	return err
}

// 关注数增加
func FollowAdd(userid int64) error {
	err := db.DB.Model(&models.User{}).Where("id = ?", userid).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error
	return err
}

// 粉丝数增加
func FollowerDown(followid int64) error {
	err := db.DB.Model(&models.User{}).Where("id = ?", followid).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error
	return err
}

// 关注数增加
func FollowDown(userid int64) error {
	err := db.DB.Model(&models.User{}).Where("id = ?", userid).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error
	return err
}

// 加入关注和粉丝列表
func Add(uerid int64, followid int64) {

}

// 移除关注和粉丝列表
func Down(uerid int64, followid int64) {

}
