package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id              int64   // 用户Id
	Password        string  `gorm:"PASSWORD()"`     // 用户密码
	FollowCount     int64   `json:"follow_count"`   // 关注数
	FollowerCount   int64   `json:"follower_count"` // 粉丝数
	IsFollow        bool    // 是否关注
	Avatar          string  // 头像
	BackgroundImage string  // 背景图
	Signature       string  // 个人简介
	TotalFavorited  int64   // 总获赞数
	WorkCount       int64   // 作品数
	FavoriteCount   int64   // 点赞数
	FollowList      []int64 //关注列表
	FollowerList    []int64 //粉丝列表
}

// 定义Follow表，记录关注的id
type FollowList struct {
	UserId   int64 `json:"user_id"`
	FollowId int64 `json:"follow_id"`
}
