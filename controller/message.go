package controller

import (
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/service"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ChatRequest struct {
	Token           string `json:"token" binding:"required"`
	ToUserId        uint   `json:"to_user_id" binding:"required"`
	LastMessageTime int64  `json:"pre_msg_time" binding:"required"`
}

type ChatResponse struct {
	StatusCode  int    `json:"status_code;string"`
	StatusMsg   string `json:"status_message"`
	MessageList []model.Message
}

type MessageRequest struct {
	Token      string `json:"token" binding:"required"`
	ToUserId   uint   `json:"to_user_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
	ActionType int    `json:"action_type" binding:"required"`
}

type MessageResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func Chat(c *gin.Context) {
	var req ChatRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, ChatResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "Argument Error",
		})
		return
	}
	claims, err := util.ParseToken(req.Token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, ChatResponse{
			StatusCode:  http.StatusUnauthorized,
			StatusMsg:   "Unauthorized User",
			MessageList: nil,
		})
		return
	}
	svc := service.NewService(c)
	list, err := svc.GetMessageList(claims.Id, req.ToUserId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, ChatResponse{
			StatusCode:  http.StatusInternalServerError,
			StatusMsg:   "Internal Error",
			MessageList: nil,
		})
		return
	}
	c.JSON(http.StatusOK, ChatResponse{
		StatusCode:  0,
		StatusMsg:   "Success",
		MessageList: list,
	})
}

func MessageAction(c *gin.Context) {
	var req MessageRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, MessageResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "参数错误",
		})
		return
	}
	claims, err := util.ParseToken(req.Token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, MessageResponse{
			StatusCode: http.StatusUnauthorized,
			StatusMsg:  "用户未认证",
		})
		return
	}
	svc := service.NewService(c)
	err = svc.SendMessage(claims.Id, req.ToUserId, req.Content, req.ActionType)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, MessageResponse{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Internal Error",
		})
		return
	}
	c.JSON(http.StatusOK, MessageResponse{
		StatusCode: 0,
		StatusMsg:  "Success",
	})
}
