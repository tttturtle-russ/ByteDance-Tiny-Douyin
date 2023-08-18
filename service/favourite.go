package service

import "ByteDance-Tiny-Douyin/model"

// 点赞的逻辑
func (svc *Service) CheckLike(userid int64, videoid int64) (bool, error) {
	like := model.FavouriteInfo{
		UserId:  userid,
		VideoId: videoid,
	}
	return svc.d.CheckLike(like)
}

func (svc *Service) RecordAdd(userid int64, videoid int64) error {
	like := model.FavouriteInfo{
		UserId:  userid,
		VideoId: videoid,
	}
	return svc.d.RecordAdd(like)
}

func (svc *Service) TotalAdd(videoid int64) error {
	like := model.VideoID{
		VideoId: videoid,
	}
	return svc.d.TotalAdd(like)
}

func (svc *Service) AuthorAdd(videoid int64) error {
	like := model.VideoID{
		VideoId: videoid,
	}
	return svc.d.AuthorAdd(like)
}

func (svc *Service) UserAdd(userid int64) error {
	like := model.UserID{
		UserId: userid,
	}
	return svc.d.UserAdd(like)
}

// 取消点赞的逻辑
func (svc *Service) RecordDelete(userid int64, videoid int64) error {
	like := model.FavouriteInfo{
		UserId:  userid,
		VideoId: videoid,
	}
	return svc.d.RecordDelete(like)
}

func (svc *Service) TotalDown(videoid int64) error {
	like := model.VideoID{
		VideoId: videoid,
	}
	return svc.d.TotalDown(like)
}

func (svc *Service) AuthorDown(videoid int64) error {
	like := model.VideoID{
		VideoId: videoid,
	}
	return svc.d.AuthorDown(like)
}

func (svc *Service) UserDown(userid int64) error {
	like := model.UserID{
		UserId: userid,
	}
	return svc.d.UserDown(like)
}

// 列表操作
func (svc *Service) LikeList(userid int64) ([]model.Video, error) {
	like := model.UserID{
		UserId: userid,
	}
	return svc.d.LikeList(like)
}
