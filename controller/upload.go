package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 实现视频上传接口
func UploadHandler(c *gin.Context) {
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
	videoURL := "/videos/" + fileName
	title := c.PostForm("title")
	authorID, _ := strconv.ParseInt(c.PostForm("author_id"), 10, 64)
	video := model.Video{CommentCount: 0, FavoriteCount: new(int64), PlayURL: videoURL, Title: title, AuthorID: authorID}
	if err = DB.Create(&video).Error; err != nil {
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
