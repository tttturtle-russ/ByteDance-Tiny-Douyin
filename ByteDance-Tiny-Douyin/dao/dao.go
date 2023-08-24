package dao

import (
	"gorm.io/gorm"
)

type Dao struct {
	*gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db.Session(&gorm.Session{}),
	}
} // NewDao 创建一个Dao实例在数据库当中创建一个新的对象
