package service

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/db"
	"github.com/gin-gonic/gin"
)

type Service struct {
	d *dao.Dao
	c *gin.Context
}

func Newservice(c *gin.Context) *Service {
	return &Service{
		c: c,
		d: dao.NewDao(db.DB),
	}
}
