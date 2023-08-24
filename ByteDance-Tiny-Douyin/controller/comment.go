package controller

import (
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/service"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type CommentRequest struct {
	Token       string
	VideoId     string
	ActionType  int
	CommentText string
	CommentID   string
}
type CommentResponse struct {
	CommentList []model.Comment
	Response
}

func CommentAction(c *gin.Context) {
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
	status, err := svc.CommentAction(req.ActionType, req.CommentID, req.CommentText, req.VideoId, time.Now())
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

func ListGet(c *gin.Context) {
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
	status, err := svc.CommentList(req.VideoId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, Response{
			StatusCode:    400,
			StatusMessage: "显示评论失败",
		})
		return
	}
	c.JSON(http.StatusOK, CommentResponse{
		Response: Response{
			StatusCode:    200,
			StatusMessage: "显示评论成功",
		},
		CommentList: status,
	})

}
