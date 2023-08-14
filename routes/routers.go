package routes

import (
	"GoWorkspace/controller"
	"GoWorkspace/midWare"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {

	//注册
	r.POST("/douyin/user/register/", controller.Register)
	//登录
	r.POST("/douyin/user/login/", controller.Login)
	r.GET("/douyin/user/", midWare.MiddleWare(), controller.Info)
	return r

}
