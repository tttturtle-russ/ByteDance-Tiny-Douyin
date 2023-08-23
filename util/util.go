package util

import (
	"ByteDance-Tiny-Douyin/controller"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var jwtKey = []byte("ByteDance-Tiny-Douyin")

type TokenClaims struct {
	Id       int64
	Username string
	jwt.StandardClaims
}

func GenerateToken(id int64, username string) (token string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	claims := TokenClaims{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ByteDance-Tiny-Douyin",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(jwtKey)
}

func ParseToken(token string) (*TokenClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*TokenClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

// GenerateUniqueFileName生成唯一文件名
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

// GetLastTime将string类型的时间转化为time.Time
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

func GenerateMassage(time time.Time) controller.VideosFeedResponse {
	var msg controller.VideosFeedResponse
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
