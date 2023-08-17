package service

import (
	"ByteDance-Tiny-Douyin/controller"
	"ByteDance-Tiny-Douyin/model"
)

func (svc *Service) CommentAction(cRequest controller.CommentRequest) (string, error) {
	if cRequest.ActionType == 1 {
		return svc.d.CommentAdd(cRequest)
	}
	return svc.d.CommentDelete(cRequest)
}
func (svc *Service) CommentList(cRequest controller.CommentRequest) ([]model.Comment, error) {
	videoId := cRequest.VideoId
	return svc.d.CommentList(videoId)
}
