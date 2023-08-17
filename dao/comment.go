package dao

import (
	"ByteDance-Tiny-Douyin/controller"
	"ByteDance-Tiny-Douyin/model"
	"time"
)

func (d *Dao) CommentAdd(request controller.CommentRequest) (string, error) {
	commentInfo := model.Comment{
		Content: request.CommentText,
		Time:    time.Now(),
		VideoID: request.VideoId,
	}
	err := d.Model(&model.Comment{}).Create(&commentInfo).Error
	if err != nil {
		return "评论失败", err
	}
	return "评论成功", nil
}
func (d *Dao) CommentDelete(request controller.CommentRequest) (string, error) {
	deleteID := request.CommentID
	err := d.Model(&model.Comment{}).Where("ID = ?", deleteID).Delete(&model.Comment{}).Error
	if err != nil {
		return "删除失败", err
	}
	return "删除成功", nil
}
func (d *Dao) CommentList(videoId string) ([]model.Comment, error) {
	videoID := videoId
	var client []model.Comment
	err := d.Model(&model.Comment{}).Where("VideoID = ?", videoID).Order("Time desc").Find(&client).Error
	if err != nil {
		return []model.Comment{}, err
	}
	return client, nil
}
