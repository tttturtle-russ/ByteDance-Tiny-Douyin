package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("ByteDance-Tiny-Douyin")

type TokenClaims struct {
	Id       uint
	Username string
	jwt.StandardClaims
}

// GenerateToken 生成带有id和username的token
func GenerateToken(id uint, username string) (token string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Hour * 24)
	claims := TokenClaims{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ByteDance-Tiny-Douyin",
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return
}

// ParseToken 解析token
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

//func SensitiveFilter(content string) bool {
//	filter := sensitive.New()
//	filter.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt")
//	validate, _ := filter.Validate(text)
//	return !validate
//}
