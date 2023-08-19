package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowListHandler(c *gin.Context) {
	//判断是否登录
	token := c.Query("token")
	if !(service.IsLogin(token)) {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 200,
			"status_msg":  "please login",
		})
	}

	//建立数据库会话
	DB := dao.NewDao(db.MySqlDB)

	var (
		videos []model.Video
		user   model.User
	)

	//获取用户id并获取其发布的所有视频
	userID := c.Query("user_id")
	err := dao.GetVideosByID(DB, userID, &videos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 400,
			"status_msg":  "get videos failed",
			"error":       err.Error(),
		})
	}

	//获取该用户所有信息
	err = dao.GetUserInfoByID(DB, userID, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 400,
			"status_msg":  "get information failed",
			"error":       err.Error(),
		})
	}

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
