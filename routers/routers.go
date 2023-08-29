package routers

import (
	"ByteDance-Tiny-Douyin/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = 50 << 20 //限制文件上传大小
	router.Static("/videos", "./videos")
	dy := router.Group("/douyin")
	dy.GET("/feed", controller.FeedHandler) // 视频流接口
	user := dy.Group("/user")
	{
		user.GET("/", controller.Userinfo)          // 用户信息
		user.POST("/register", controller.Register) // 注册
		user.POST("/login", controller.Login)       // 登录
	}
	favorite := dy.Group("/favorite")
	{
		favorite.POST("/action", controller.FavouriteAction) // 点赞和取消点赞
		favorite.GET("/list", controller.FavouriteList)      // 点赞列表
	}
	publish := dy.Group("/publish")
	{
		publish.POST("/action", controller.UploadHandler) // 用户上传视频
		publish.GET("/list", controller.ShowListHandler)  // 视频列表
	}
	comment := dy.Group("/comment")
	{
		comment.POST("/action", controller.CommentAction) // 评论和回复
		comment.GET("/list", controller.ListGet)          // 评论列表
	}
	relation := dy.Group("/relation")
	{
		relation.POST("/action")                            // 关注和取消关注
		relation.GET("/follow/list")                        // 关注列表
		relation.GET("/follower/list")                      // 粉丝列表
		relation.GET("/friend/list", controller.FriendList) // 好友列表
	}
	message := dy.Group("/message")
	{
		message.GET("/chat", controller.Chat)             // 用户的聊天消息记录
		message.POST("/action", controller.MessageAction) // 发送消息
	}
	return router
}
