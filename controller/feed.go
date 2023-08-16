package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FeedHandler(c *gin.Context) {
	const videoMaxNum int = 10
	var (
		msg    model.MessageReturned
		videos []model.Video
		user   model.User
	)

	//获取latest_time参数并将其转换为time.Time类型
	latestTime, err := service.GetLastTime(c.Query("latest_time"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	//按ID降序获得最多10条数据
	err = dao.DB.Where("created_at < ?", latestTime).Order("created_at desc").Limit(videoMaxNum).Find(&videos).Error
	if err != nil {
		msg = service.GenerateMassage(latestTime)
		c.JSON(http.StatusOK, msg)
		return
	}

	//遍历videos切片将author字段绑定
	for index, video := range videos {
		err = dao.DB.Where("id = ?", video.AuthorID).First(&user).Error
		if err != nil {
			msg = service.GenerateMassage(latestTime)
			c.JSON(http.StatusOK, msg)
			return
		}
		videos[index].Author = user
	}

	//绑定msg
	msg.VideoList = videos
	//将最早的视频的created_at转换为时间戳返回
	msg.NextTime = videos[len(videos)-1].CreatedAt.Unix()
	msg.StatusMsg = "get succeed"
	msg.StatusCode = 0

	//返回响应
	c.JSON(http.StatusOK, msg)
}
