package controller

import (
	"ByteDance-Tiny-Douyin/service"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CommentRequest struct {
	Token       string
	VideoId     string
	ActionType  int
	CommentText string
	CommentID   string
}

func commentAction(c *gin.Context) {
	var req CommentRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response{
			StatusCode:    400,
			StatusMessage: "参数错误",
		})
		return
	}
	_, err = util.ParseToken(req.Token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, Response{
			StatusCode:    401,
			StatusMessage: "用户未认证",
		})
	}
	svc := service.NewService(c)
	status, err := svc.Comment(req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response{
			StatusCode:    400,
			StatusMessage: status,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode:    200,
		StatusMessage: status,
	})
}
