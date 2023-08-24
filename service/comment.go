package service

import (
	"ByteDance-Tiny-Douyin/model"
	"time"
)

func (svc *Service) CommentAction(ActionType int, CommentID string, CommentText string, videoID string, time time.Time) (string, error) {
	if ActionType == 1 {
		return svc.d.CommentAdd(CommentText, time, videoID)
	}
	return svc.d.CommentDelete(CommentID)
}
func (svc *Service) CommentList(VideoID string) ([]model.Comment, error) {
	return svc.d.CommentList(VideoID)
}
