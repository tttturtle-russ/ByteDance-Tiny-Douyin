package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideosListRequest struct {
	Token  string `json:"token" binding:"required"`   // 用户鉴权token
	UserID string `json:"user_id" binding:"required"` // 用户id
}

type VideosListResponse struct {
	StatusCode int64         `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `json:"status_msg"`  // 返回状态描述
	VideoList  []model.Video `json:"video_list"`  // 用户发布的视频列表
}

func ShowListHandler(c *gin.Context) {
	//判断是否登录
	token := c.Query("token")
	if !(util.IsLogin(token)) {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 200,
			"status_msg":  "please login",
		})
		return
	}

	//建立数据库会话
	DB := dao.NewDao(db.MySqlDB)
	videos := []model.Video{}
	user := model.User{}

	//获取用户id并获取其发布的所有视频
	userID := c.Query("user_id")
	err := dao.GetVideosByID(DB, userID, &videos)
	if err != nil {
		c.JSON(http.StatusBadRequest, VideosListResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "get user_id failed",
		})
		return
	}

	//获取该用户所有信息
	err = dao.GetUserInfoByID(DB, userID, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, VideosListResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "get user_information failed",
		})
	}

	//将videos.author与用户信息绑定
	for index, _ := range videos {
		videos[index].Author = user
	}

	c.JSON(http.StatusOK, VideosListResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "get video list succeed",
		VideoList:  videos,
	})
}
