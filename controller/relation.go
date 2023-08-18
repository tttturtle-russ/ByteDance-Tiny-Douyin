package controller

import (
	"ByteDance-Tiny-Douyin/model"
	"github.com/gin-gonic/gin"
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

}
