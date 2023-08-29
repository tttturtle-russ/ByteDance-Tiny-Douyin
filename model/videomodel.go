package model

import (
	"gorm.io/gorm"
)

//用来定义视频流接口所需的model

// Video用来存储视频相关的信息
type Video struct {
	gorm.Model
	Author        User   `gorm:"-"`              // 视频作者信息
	AuthorID      int64  `json:"author_id"`      //视频作者ID
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount *int64 `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}
