package dao

import (
	"ByteDance-Tiny-Douyin/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type Dao struct {
	*gorm.DB
}

// NewDao 创建一个Dao实例
func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db.Session(&gorm.Session{}),
	}
}

// UploadVideo 将上传的video信息插入到数据库
func UploadVideo(c *gin.Context, fileName string, DB *Dao) (err error) {
	videoURL := "/videos/" + fileName
	title := c.PostForm("title")
	authorID, _ := strconv.ParseInt(c.PostForm("author_id"), 10, 64)
	video := model.Video{CommentCount: 0, FavoriteCount: new(int64), PlayURL: videoURL, Title: title, AuthorID: authorID}
	err = DB.Create(&video).Error
	return err
}

// GetVideosByID 根据传入的用户ID获取该用户发布的所有视频信息
func GetVideosByID(DB *Dao, userID string, videos *[]model.Video) (err error) {
	err = DB.Where("author_id = ?", userID).Find(videos).Error
	return err
}

// GetUserInfoByID 根据传入的用户ID获取该用户信息
func GetUserInfoByID(DB *Dao, userID string, user *model.User) (err error) {
	err = DB.Where("id = ?", userID).First(&user).Error
	return err
}
