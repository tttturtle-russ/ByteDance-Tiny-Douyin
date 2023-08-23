package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UploadResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// 实现视频上传接口
func UploadHandler(c *gin.Context) {
	//判断是否登录
	token := c.PostForm("token")
	if !(util.IsLogin(token)) {
		c.JSON(http.StatusBadRequest, UploadResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "please login",
		})
		return
	}

	//建立数据库会话
	DB := dao.NewDao(db.MySqlDB)

	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusBadRequest, UploadResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "upload file failed",
		})
		return
	}
	//保存文件到指定目录
	fileName := util.GenerateUniqueFileName(file.Filename)
	filePath := "videos/" + fileName
	if err = c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusBadRequest, UploadResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "save file failed",
		})
		return
	}
	//将信息存入数据库
	if err = dao.UploadVideo(c, fileName, DB); err != nil {
		c.JSON(http.StatusBadRequest, UploadResponse{
			StatusCode: http.StatusBadRequest,
			StatusMsg:  "save data failed",
		})
		return
	}
	//返回状态码状态信息
	c.JSON(http.StatusOK, UploadResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "upload file succeed",
	})
	return
}
