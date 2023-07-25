package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()
	dy := router.Group("/douyin")
	dy.GET("/feed") // 视频流接口
	user := dy.Group("/user")
	{
		user.GET("/")          // 用户信息
		user.POST("/register") // 注册
		user.POST("/login")    // 登录
	}
	favorite := dy.Group("/favorite")
	{
		favorite.POST("/action") // 点赞和取消点赞
		favorite.GET("/list")    // 点赞列表
	}
	publish := dy.Group("/publish")
	{
		publish.POST("/action") // 用户上传视频
		publish.GET("/list")    // 视频列表
	}
	comment := dy.Group("/comment")
	{
		comment.POST("/action") // 评论和回复
		comment.GET("/list")    // 评论列表
	}
	relation := dy.Group("/relation")
	{
		relation.POST("/action")       // 关注和取消关注
		relation.GET("/follow/list")   // 关注列表
		relation.GET("/follower/list") // 粉丝列表
		relation.GET("/friend/list")   // 好友列表
	}
	message := dy.Group("/message")
	{
		message.GET("/chat")    // 用户的聊天消息记录
		message.POST("/action") // 发送消息
	}
	return router
}
