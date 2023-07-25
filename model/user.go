package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string // 用户名
	Password        string // 用户密码
	FollowCount     int64  // 关注数
	FollowerCount   int64  // 粉丝数
	IsFollow        bool   // 是否关注
	Avatar          string // 头像
	BackgroundImage string // 背景图
	Signature       string // 个人简介
	TotalFavorited  int64  // 总获赞数
	WorkCount       int64  // 作品数
	FavoriteCount   int64  // 点赞数
}
