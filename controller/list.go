package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowListHandler(c *gin.Context) {
	var (
		videos []model.Video
		user   model.User
		msg    model.MessageReturned
	)

	//获取用户id并获取其发布的所有视频
	userID := c.Query("user_id")
	dao.DB.Where("author_id = ?", userID).Find(&videos)

	//获取该用户所有信息
	dao.DB.Where("id = ?", userID).First(&user)

	//将videos.author与用户信息绑定
	for index, _ := range videos {
		videos[index].Author = user
	}

	msg.VideoList = videos
	msg.StatusCode = http.StatusOK
	msg.StatusMsg = "succeed"

	c.JSON(http.StatusOK, msg)
}
