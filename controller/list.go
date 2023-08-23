package controller

import (
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/service"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	judge, err := util.IsLogin(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, UploadResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "token解析错误",
		})
	}
	if !judge {
		c.JSON(http.StatusBadRequest, VideosListResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "please login",
		})
		return
	}

	//初始化一个user 和video实例
	videos := []model.Video{}
	user := model.User{}

	//绑定参数
	var req VideosListRequest
	err = c.ShouldBindQuery(&req)
	if err != nil {
		log.Printf("绑定参数失败")
		c.JSON(http.StatusBadRequest, UploadResponse{
			StatusCode: 400,
			StatusMsg:  err.Error(),
		})
		return
	}

	//获取用户id并获取其发布的所有视频
	svc := service.NewService(c)
	videos, err = svc.GetVideosByID(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, VideosListResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "get user_id failed",
		})
		return
	}

	//获取该用户所有信息
	userID, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		log.Println("获取用户ID失败")
		c.JSON(http.StatusBadRequest, VideosListResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "获取用户ID失败",
		})
		return
	}
	user, err = svc.GetInfoByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, VideosListResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "get user_information failed",
		})
		return
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
