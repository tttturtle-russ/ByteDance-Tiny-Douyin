package dao

import "ByteDance-Tiny-Douyin/model"

// GetVideosByID 根据传入的用户ID获取该用户发布的所有视频信息
func (d *Dao) GetVideosByID(userID string) (videos []model.Video, err error) {
	err = d.Where("author_id = ?", userID).Find(videos).Error
	return videos, err
}
