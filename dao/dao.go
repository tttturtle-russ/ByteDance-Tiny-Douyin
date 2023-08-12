package dao

import "gorm.io/gorm"

type Dao struct {
	*gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db.Session(&gorm.Session{}),
	}
}
