package dao

import (
	"ByteDance-Tiny-Douyin/model"
	"golang.org/x/crypto/bcrypt"
)

func (d *Dao) RegisterUser(user model.User) (int64, error) {
	err := d.Model(&model.User{}).Create(&user).Error
	if err != nil {
		return 0, err
	}
	//err = d.Model(user).Where(&user).Select("id").Scan(&id).Error
	return user.ID, err
}
func (d *Dao) LoginUser(user model.User) (int64, error) {
	name := user.Name
	password := user.Password
	client := model.User{}
	err := d.Model(&model.User{}).Where("Name = ?", name).First(&client).Error
	if client.ID == 0 || err != nil {
		return 0, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password)); err != nil {
		return 0, err
	}

	return user.ID, err
}

// GetUserInfoByID 根据传入的用户ID获取该用户信息
func GetUserInfoByID(DB *Dao, userID string, user *model.User) (err error) {
	err = DB.Where("id = ?", userID).First(&user).Error
	return err
}
