package db

import (
	"GoWorkspace/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var realdb *gorm.DB

var EncryptionKey = []byte("token")

type Claims struct {
	UserID uint
	jwt.StandardClaims
}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "user"
	username := "root"
	password := "oyyz20031010"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	tempdb, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	tempdb.AutoMigrate(&model.User{})
	realdb = tempdb
	return tempdb

}

func GetDB() *gorm.DB {
	return realdb
}

func SendToken(user model.User) (string, error) {
	Time := time.Now().Add(24 * time.Hour)

	claim := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: Time.Unix(),       //过期时间
			IssuedAt:  time.Now().Unix(), //当前时间
			Issuer:    "127.0.0.1",
			Subject:   "token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(EncryptionKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return EncryptionKey, nil
	})
	return token, claims, err
}
