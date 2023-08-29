package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	Content string
	Time    time.Time
	VideoID string
}
