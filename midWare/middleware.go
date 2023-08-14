package midWare

import (
	"GoWorkspace/db"
	"GoWorkspace/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func MiddleWare() gin.HandlerFunc { //提前过滤一遍用户是否合法
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		//	有效token的验证,若token为空或者前缀为Bearer说明并非有效
		if tokenString == "" || strings.HasPrefix(tokenString, "Bearer") {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不够"})
			context.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := db.ParseToken(tokenString)
		if err != nil || token.Valid == false {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不够"})
			context.Abort()
		} //验证token是否合法有效
		userid := claims.UserID
		realdb := db.GetDB()
		var user model.User
		realdb.First(&user, userid)

		if user.ID == 0 {
			context.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不够"})
			context.Abort()
			return
		} //没有查询到用户信息
		context.Set("user", user) //写入上下文方便后续继续调用
		context.Next()
	}
}
