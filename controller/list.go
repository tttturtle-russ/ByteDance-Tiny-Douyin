package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowListHandler(c *gin.Context) {
	//建立数据库会话
	DB := dao.NewDao(db.MySqlDB)

	var (
		videos []model.Video
		user   model.User
	)

	//获取用户id并获取其发布的所有视频
	userID := c.Query("user_id")
	DB.Where("author_id = ?", userID).Find(&videos)

	//获取该用户所有信息
	DB.Where("id = ?", userID).First(&user)

	//将videos.author与用户信息绑定
	for index, _ := range videos {
		videos[index].Author = user
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "succeed",
		"video_list":  videos,
	})
}
