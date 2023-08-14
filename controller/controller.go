package controller

import (
	"GoWorkspace/db"
	"GoWorkspace/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(ctx *gin.Context) {

	realdb := db.GetDB()

	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	RegisterUser := model.User{}
	ctx.Bind(&RegisterUser)
	name := RegisterUser.Name
	password := RegisterUser.Name

	//数据验证
	if len(name) == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名不能为空",
		})
		return
	}
	if len(password) < 8 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码不得少于8位，请重新注册",
		})
		return
	}

	//判断手机号是否存在
	var repeatUser model.User
	realdb.Where("name = ?", name).First(&repeatUser)
	if repeatUser.ID != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户已存在,请不要重复注册",
		})
		return
	}
	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	newUser := model.User{
		Name:      name,
		Password:  string(hasedPassword),
		Subscribe: 0,
		Fans:      0,
	}
	realdb.Create(&newUser)

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

func Login(ctx *gin.Context) {

	realdb := db.GetDB()

	//获取参数
	//此处使用Bind()函数，可以处理不同格式的前端数据
	var loginUser model.User
	ctx.Bind(&loginUser)
	name := loginUser.Name
	password := loginUser.Password
	client := model.User{}
	realdb.Where("name = ?", name).First(&client)
	if client.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名或者密码错误，请核验后重新登陆",
		})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "用户名或者密码错误，请核验后重新登陆",
		})
		return
	}
	token, err := db.SendToken(client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "token获取失败",
		})
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    gin.H{"token": token},
		"message": "登录成功",
	})
}

func Info(context *gin.Context) {
	user, _ := context.Get("user")
	context.JSON(http.StatusOK, gin.H{
		"data": gin.H{"user": user},
	})

}
