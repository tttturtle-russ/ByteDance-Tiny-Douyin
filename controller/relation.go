package controller

import (
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/service"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type FriendListRequest struct {
	UserId int64  `json:"user_id" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

type FriendListResponse struct {
	StatusCode int          `json:"status_code;string"`
	StatusMsg  string       `json:"status_msg"`
	Friends    []model.User `json:"user_list"`
}

func FriendList(c *gin.Context) {
	var req FriendListRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, FriendListResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "Argument Error",
			Friends:    nil,
		})
		return
	}
	claims, err := util.ParseToken(req.Token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, FriendListResponse{
			StatusCode: http.StatusUnauthorized,
			StatusMsg:  "Unauthorized User",
			Friends:    nil,
		})
		return
	}
	svc := service.NewService(c)
	list, err := svc.GetFriendList(claims.Id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, FriendListResponse{
			StatusCode: http.StatusInternalServerError,
			StatusMsg:  "Internal Server Error",
			Friends:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, FriendListResponse{
		StatusCode: 0,
		StatusMsg:  "Success",
		Friends:    list,
	})
}
