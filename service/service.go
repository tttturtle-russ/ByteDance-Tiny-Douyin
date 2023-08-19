package service

import (
	"ByteDance-Tiny-Douyin/model"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func GenerateUniqueFileName(originalName string) string {
	// 获取当前时间戳
	timestamp := time.Now().Unix()

	// 生成随机字符串
	rand.Seed(time.Now().UnixNano())
	randomString := strconv.FormatInt(rand.Int63(), 36)

	// 组合生成唯一文件名
	uniqueName := fmt.Sprintf("%d_%s_%s", timestamp, randomString, originalName)
	return uniqueName
}

func GetLastTime(nowTime string) (latestTime time.Time, err error) {
	if nowTime == "" {
		latestTime = time.Unix(time.Now().Unix(), 0)
	} else {
		tmpTime, err := strconv.ParseInt(nowTime, 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		latestTime = time.Unix(tmpTime, 0)
	}
	return latestTime, nil
}

func GenerateMassage(time time.Time) model.MessageReturned {
	var msg model.MessageReturned
	msg.StatusMsg = "get failed"
	msg.VideoList = nil
	msg.NextTime = time.Unix()
	msg.StatusCode = http.StatusOK
	return msg
}

func IsLogin(token string) bool {
	if token == "" {
		return false
	}
	return true
}
