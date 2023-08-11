package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func FeedHandler(c *gin.Context) {
	const videoMaxNum int = 10
	var msg models.MessageReturned
	var videos []models.Video
	var user models.User
	msg.StatusMsg = new(string)
	msg.NextTime = new(int64)

	//获取latest_time参数并将其转换为time.Time类型
	var latestTime time.Time
	//特判latest_time为空时 默认以当前时间
	if a := c.Query("latest_time"); a == "" {
		latestTime = time.Unix(time.Now().Unix(), 0)
	} else {
		tmpTime, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			panic(err.Error())
		}
		latestTime = time.Unix(tmpTime, 0)
	}

	//按ID降序获得最多10条数据
	err := dao.DB.Where("created_at < ?", latestTime).Order("created_at desc").Limit(videoMaxNum).Find(&videos).Error
	if err != nil {
		*msg.StatusMsg = "get failed"
		msg.VideoList = nil
		*msg.NextTime = latestTime.Unix()
		msg.StatusCode = http.StatusOK
		c.JSON(http.StatusOK, msg)
		return
	}

	//遍历videos切片将author字段绑定
	for index, video := range videos {
		err = dao.DB.Where("id = ?", video.AuthorID).First(&user).Error
		if err != nil {
			*msg.StatusMsg = "get failed"
			msg.VideoList = nil
			*msg.NextTime = latestTime.Unix()
			msg.StatusCode = http.StatusOK
			c.JSON(http.StatusOK, msg)
			return
		}
		videos[index].Author = user
	}

	//绑定msg
	msg.VideoList = videos
	//将最早的视频的created_at转换为时间戳返回
	*msg.NextTime = videos[len(videos)-1].CreatedAt.Unix()
	*msg.StatusMsg = "get succeed"
	msg.StatusCode = http.StatusOK

	//返回响应
	c.JSON(http.StatusOK, msg)
}
