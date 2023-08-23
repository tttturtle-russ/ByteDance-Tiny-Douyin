package controller

import (
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/service"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type VideosFeedRequest struct {
	LatestTime string `json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `json:"token,omitempty"`       // 用户登录状态下设置
}

type VideosFeedResponse struct {
	NextTime   int64         `json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int64         `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `json:"status_msg"`  // 返回状态描述
	VideoList  []model.Video `json:"video_list"`  // 视频列表
}

func FeedHandler(c *gin.Context) {
	var req VideosFeedRequest
	msg := VideosFeedResponse{}
	videos := []model.Video{}
	user := model.User{}
	svc := service.NewService(c)

	//获取latest_time参数并将其转换为time.Time类型
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Println()
		c.JSON(http.StatusBadRequest, VideosFeedResponse{
			StatusMsg:  "Argument error",
			StatusCode: http.StatusOK,
		})
	}
	latestTime, err := util.GetLastTime(req.LatestTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, VideosFeedResponse{
			StatusMsg:  "get latest_time failed",
			StatusCode: http.StatusOK,
		})
	}

	//按ID降序获得最多10条数据
	videos, err = svc.GetVideoByTime(latestTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, VideosFeedResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "get videos failed",
		})
		return
	}

	//遍历videos切片将author字段绑定
	for index, video := range videos {
		user, err = svc.GetInfoByUserID(video.AuthorID)
		if err != nil {
			c.JSON(http.StatusBadRequest, VideosFeedResponse{
				StatusCode: http.StatusBadRequest,
				StatusMsg:  "get author_info failed",
			})
			return
		}
		videos[index].Author = user
	}

	//绑定msg
	msg.VideoList = videos
	msg.NextTime = videos[len(videos)-1].CreatedAt.Unix()
	msg.StatusMsg = "get succeed"
	msg.StatusCode = 0

	//返回响应
	c.JSON(http.StatusOK, msg)
}
