package dao

import (
	"ByteDance-Tiny-Douyin/model"
)

// UploadVideo 将上传的video信息插入到数据库
func (d *Dao) UploadVideo(fileName string, title string, authorID int64) (err error) {
	videoURL := "/videos/" + fileName
	video := model.Video{CommentCount: 0, FavoriteCount: new(int64), PlayURL: videoURL, Title: title, AuthorID: authorID}
	err = d.Create(&video).Error
	return err
}
