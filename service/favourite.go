package service

import "ByteDance-Tiny-Douyin/model"

// 点赞的逻辑
func (svc *Service) CheckLikeList(userid int64, videoid int64) (bool, error) {
	like := model.FavouriteInfo{
		UserId:  userid,
		VideoId: videoid,
	}
	return svc.d.CheckLikeList(like)
}

func (svc *Service) LikeListAdd(userid int64, videoid int64) error {
	like := model.FavouriteInfo{
		UserId:  userid,
		VideoId: videoid,
	}
	return svc.d.LikeListAdd(like)
}

func (svc *Service) VideoFavoriteCountAdd(videoid int64) error {
	like := model.VideoID{
		VideoId: videoid,
	}
	return svc.d.VideoFavoriteCountAdd(like)
}

func (svc *Service) UserTotalFavoritedAdd(videoid int64) error {
	like := model.VideoID{
		VideoId: videoid,
	}
	return svc.d.UserTotalFavoritedAdd(like)
}

func (svc *Service) UserFavoriteCountAdd(userid int64) error {
	like := model.UserID{
		UserId: userid,
	}
	return svc.d.UserFavoriteCountAdd(like)
}

// 取消点赞的逻辑
func (svc *Service) LikeListDelete(userid int64, videoid int64) error {
	like := model.FavouriteInfo{
		UserId:  userid,
		VideoId: videoid,
	}
	return svc.d.LikeListDelete(like)
}

func (svc *Service) VideoFavoriteCountDown(videoid int64) error {
	like := model.VideoID{
		VideoId: videoid,
	}
	return svc.d.VideoFavoriteCountDown(like)
}

func (svc *Service) UserTotalFavoritedDown(videoid int64) error {
	like := model.VideoID{
		VideoId: videoid,
	}
	return svc.d.UserTotalFavoritedDown(like)
}

func (svc *Service) UserFavoriteCountDown(userid int64) error {
	like := model.UserID{
		UserId: userid,
	}
	return svc.d.UserFavoriteCountDown(like)
}

// 列表操作
func (svc *Service) FindVideosInLikeList(userid int64) ([]model.Video, error) {
	like := model.UserID{
		UserId: userid,
	}
	return svc.d.FindVideosInLikeList(like)
}
