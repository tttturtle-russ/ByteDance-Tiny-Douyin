package service

import "ByteDance-Tiny-Douyin/controller"

func (svc *Service) Comment(cRequest controller.CommentRequest) (string, error) {
	if cRequest.ActionType == 1 {
		return svc.d.CommentAdd(cRequest)
	}
	return svc.d.CommentDelete(cRequest)
}
