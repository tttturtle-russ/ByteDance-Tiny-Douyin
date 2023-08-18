package dao

import (
	"ByteDance-Tiny-Douyin/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
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

// GetVideoByTime 根据视频发布时间获取视频信息
func GetVideoByTime(DB *Dao, latestTime time.Time, videos *[]model.Video) (err error) {
	const videoMaxNum int = 10
	err = DB.Where("created_at < ?", latestTime).Order("created_at desc").Limit(videoMaxNum).Find(videos).Error
	return err
}
