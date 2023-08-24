package controller

import (
	"ByteDance-Tiny-Douyin/service"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type needResponse struct {
	Response
	Id    uint   `json:"user_id"`
	Token string `json:"token"`
}

type userRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Register(c *gin.Context) {
	// 先进行参数校验和绑定
	var req userRequest
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
	c.JSON(http.StatusOK, needResponse{
		Response: Response{
			StatusCode:    200,
			StatusMessage: "注册成功",
		},
		Id:    id,
		Token: "",
	})
}
func Login(c *gin.Context) {
	var loginRequest userRequest
	err := c.ShouldBindQuery(&loginRequest)
	if err != nil {
		log.Printf("绑定参数失败")
		c.JSON(http.StatusBadRequest, Response{
			StatusCode:    400,
			StatusMessage: err.Error(),
		})
		return
	}
	svc := service.NewService(c)
	id, err := svc.LoginUser(loginRequest.Username, loginRequest.Password)
	token, err := util.GenerateToken(id, loginRequest.Username)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, Response{
			StatusCode:    400,
			StatusMessage: err.Error(),
		})
		return
	} else if id == 0 && token == "" {
		c.JSON(http.StatusUnprocessableEntity, Response{
			StatusCode:    422,
			StatusMessage: "用户名或者密码错误，请重试！",
		})
		return
	}
	// 返回给前端结果
	c.JSON(http.StatusOK, needResponse{
		Response: Response{
			StatusCode:    200,
			StatusMessage: "登陆成功",
		},
		Id:    id,
		Token: token,
	})
}
