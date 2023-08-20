package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 实现视频上传接口
func UploadHandler(c *gin.Context) {
	//判断是否登录
	token := c.Query("token")
	if !(service.IsLogin(token)) {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 200,
			"status_msg":  "please login",
		})
		return
	}

	//建立数据库会话
	DB := dao.NewDao(db.MySqlDB)

	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//保存文件到指定目录
	fileName := service.GenerateUniqueFileName(file.Filename)
	filePath := "videos/" + fileName
	if err = c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//将信息存入数据库
	if err = dao.UploadVideo(c, fileName, DB); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	//返回状态码状态信息
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "ok",
	})
}
