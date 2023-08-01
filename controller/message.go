package controller

import (
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/service"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type MessageChatRequest struct {
	Token           string `json:"token" binding:"required"`
	ToUserId        int64  `json:"to_user_id" binding:"required"`
	LastMessageTime int64  `json:"pre_msg_time" binding:"required"`
}

type MessageChatResponse struct {
	Response
	MessageList []model.Message
}

func Chat(c *gin.Context) {
	var req MessageChatRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, MessageChatResponse{
			Response: Response{
				StatusCode:    400,
				StatusMessage: "参数错误",
			},
		})
		return
	}
	claims, err := util.ParseToken(req.Token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, MessageChatResponse{
			Response: Response{
				StatusCode:    401,
				StatusMessage: "用户未认证",
			},
			MessageList: nil,
		})
		return
	}
	svc := service.NewService(c)
	list, err := svc.GetMessageList(claims.Id, req.ToUserId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, MessageChatResponse{
			Response: Response{
				StatusCode:    500,
				StatusMessage: "服务器错误",
			},
			MessageList: nil,
		})
		return
	}
	c.JSON(http.StatusOK, MessageChatResponse{
		Response: Response{
			StatusCode:    200,
			StatusMessage: "获取成功",
		},
		MessageList: list,
	})
}

func MessageAction(c *gin.Context) {

}
