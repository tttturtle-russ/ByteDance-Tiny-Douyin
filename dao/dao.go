package dao

import "gorm.io/gorm"

type Dao struct {
	*gorm.DB
}

// dao实例
func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db.Session(&gorm.Session{}),
	}
}
