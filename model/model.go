package model

// 定义Like表，记录点赞的id和对应的videoid
type LikeList struct {
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

// service
type FavouriteInfo struct {
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

type VideoID struct {
	VideoId int64 `json:"video_id"`
}

type UserID struct {
	UserId int64 `json:"user_id"`
}
