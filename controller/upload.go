package controller

import (
	"ByteDance-Tiny-Douyin/service"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UploadRequest struct {
	Token string `form:"token" binding:"required"`
	Title string `form:"title" binding:"required"`
}

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

	//绑定参数
	var req UploadRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		log.Printf("绑定参数失败")
		c.JSON(http.StatusBadRequest, UploadResponse{
			StatusCode: 400,
			StatusMsg:  err.Error(),
		})
		return
	}

	//解析token获得userID
	claims, err := util.ParseToken(req.Token)
	if err != nil {
		log.Printf("获取ID失败")
		c.JSON(http.StatusBadRequest, UploadResponse{
			StatusCode: 400,
			StatusMsg:  err.Error(),
		})
		return
	}
	authorID := claims.Id

	//获取视频数据
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
	var svc service.Service
	if err = svc.UploadVideo(fileName, req.Title, authorID); err != nil {
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
