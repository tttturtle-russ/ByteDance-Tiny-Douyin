package controller

import (
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/service"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type T struct {
}
type userinfoRequest struct {
	Token string `json:"token"binding:"required"`
}
type userinfoResponse struct {
	Response
	userinfo model.User
}

func Userinfo(c *gin.Context) {
	var req userinfoRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, userinfoResponse{
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
		c.JSON(http.StatusUnauthorized, userinfoResponse{
			Response: Response{
				StatusCode:    401,
				StatusMessage: "用户未认证",
			},
			userinfo: model.User{},
		})

	}

	svc := service.NewService(c)
	list, err := svc.GetuserInfo(claims.Username)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, userinfoResponse{
			Response: Response{
				StatusCode:    500,
				StatusMessage: "服务器错误",
			},
			userinfo: model.User{},
		})
		return
	}
	c.JSON(http.StatusOK, userinfoResponse{
		Response: Response{
			StatusCode:    200,
			StatusMessage: "获取成功",
		},
		userinfo: list,
	})
}
