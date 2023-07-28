package controller

import (
	"ByteDance-Tiny-Douyin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RegisterResponse struct {
	Response
	Id    uint   `json:"user_id"`
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Register(c *gin.Context) {
	// 先进行参数校验和绑定
	var req RegisterRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Printf("绑定参数失败")
		c.JSON(http.StatusBadRequest, Response{
			StatusCode:    400,
			StatusMessage: err.Error(),
		})
		return
	}
	// 创建一个新的service,用于处理业务
	svc := service.NewService(c)
	id, err := svc.RegisterUser(req.Username, req.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, Response{
			StatusCode:    400,
			StatusMessage: err.Error(),
		})
		return
	}
	// 返回给前端结果
	c.JSON(http.StatusOK, RegisterResponse{
		Response: Response{
			StatusCode:    200,
			StatusMessage: "注册成功",
		},
		Id:    id,
		Token: "",
	})
}
