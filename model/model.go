package model

type User struct {
	Id              int64
	Name            string // 用户名
	Password        string `gorm:"PASSWORD()"` // 用户密码
	FollowCount     int64  // 关注数
	FollowerCount   int64  // 粉丝数
	IsFollow        bool   // 是否关注
	Avatar          string // 头像
	BackgroundImage string // 背景图
	Signature       string // 个人简介
	TotalFavorited  int64  `json:"total_favorited"` // 总获赞数
	WorkCount       int64  // 作品数
	FavoriteCount   int64  // 点赞数
	//VideoFavourites []Video `json:"video_favourites"` //点赞的视频id
}

type Video struct {
	Author        User   `gorm:"-"`              // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
	//UserFavourites []int64 `json:"user_favourites"` //点赞的用户id
}

// 定义Like表，记录点赞的id和对应的videoid
type LikeList struct {
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
}

// 常用到userid和videoid的结构体
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
