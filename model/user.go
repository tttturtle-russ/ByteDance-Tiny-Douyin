package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name            string // 用户名
	Password        string `gorm:"PASSWORD()"` // 用户密码
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

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		var password []byte
		password, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		tx = tx.Model(&User{}).Where("name = ?", u.Name)
		tx.Update("password", password)
		tx.Select("id").Scan(&u.ID)
	}
	return
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if err := tx.Model(&User{}).Where("name = ?", u.Name).First(&User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.New("用户已存在")
}
