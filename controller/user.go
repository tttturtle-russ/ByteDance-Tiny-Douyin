package controller

import (
	"ByteDance-Tiny-Douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterResponse struct {
	Response
	Id    int64  `json:"user_id"`
	Token string `json:"token"`
}

func Register(c *gin.Context) {
	//username := c.Query("username")
	//password := c.Query("password")
	type Argument struct {
		UserName string `form:"username" binding:"required"`
		Password string `form:"password" binding:"required"`
	}
	a := Argument{}
	err := c.ShouldBindQuery(&a)
	err = service.RegisterUser(a.UserName, a.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "注册失败",
		})
		return
	}
	c.JSON(http.StatusOK, RegisterResponse{
		Response: Response{
			StatusCode:    http.StatusOK,
			StatusMessage: "注册成功",
		},
		Id:    0,
		Token: "",
	})
}
