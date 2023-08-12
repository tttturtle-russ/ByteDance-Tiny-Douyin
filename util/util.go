package util

import (
	"github.com/dgrijalva/jwt-go"
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
