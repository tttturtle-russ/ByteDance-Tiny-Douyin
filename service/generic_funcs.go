package service

import (
	"fmt"
	"math/rand"
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
