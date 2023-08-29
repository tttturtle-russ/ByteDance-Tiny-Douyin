package dao

import (
	"ByteDance-Tiny-Douyin/model"
	"gorm.io/gorm"
)

//点赞

// 查找LikeList里面是否已存在记录
func (d *Dao) CheckLikeList(LikeInfo model.FavouriteInfo) (bool, error) {
	var like model.LikeList
	userid := LikeInfo.UserId
	videoid := LikeInfo.VideoId

	err := d.Model(&model.LikeList{}).Where("user_id = ? AND video_id = ?", userid, videoid).First(&like)

	if err != nil {
		return false, err.Error //查询记录不存在或查询出错
	} else {
		return true, err.Error
	}
}

// 增加LikeList记录
func (d *Dao) LikeListAdd(LikeInfo model.FavouriteInfo) error {
	userid := LikeInfo.UserId
	videoid := LikeInfo.VideoId

	add := &model.LikeList{
		UserId:  userid,
		VideoId: videoid,
	}

	result := d.Model(&model.LikeList{}).Create(add)
	return result.Error
}

// video对应的总获赞数
func (d *Dao) VideoFavoriteCountAdd(VideoInfo model.VideoID) error {
	videoid := VideoInfo.VideoId

	err := d.Model(&model.Video{}).Where("id = ?", videoid).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	return err
}

// video作者的获赞数
func (d *Dao) UserTotalFavoritedAdd(VideoInfo model.VideoID) error {
	var video model.Video
	videoid := VideoInfo.VideoId
	err := d.Model(&model.Video{}).Where("id = ?", videoid).First(&video).Error
	if err != nil {
		return err
	}

	authorid := video.Author.Id
	err = d.Model(&model.User{}).Where("id = ?", authorid).UpdateColumn("total_favorited", gorm.Expr("total_favorited + ?", 1)).Error
	if err != nil {
		return err
	}
	return err
}

// 用户的点赞数
func (d *Dao) UserFavoriteCountAdd(UserInfo model.UserID) error {
	userid := UserInfo.UserId
	err := d.Model(&model.User{}).Where("id = ?", userid).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error
	return err
}

//取消点赞

// 删除LikeList记录
func (d *Dao) LikeListDelete(LikeInfo model.FavouriteInfo) error {
	userid := LikeInfo.UserId
	videoid := LikeInfo.VideoId
	result := d.Model(&model.LikeList{}).Where("user_id = ? AND video_id = ?", userid, videoid).Delete(&model.LikeList{})
	return result.Error
}

// 减少video总点赞数
func (d *Dao) VideoFavoriteCountDown(VideoInfo model.VideoID) error {
	videoid := VideoInfo.VideoId
	err := d.Model(&model.Video{}).Where("id = ?", videoid).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	return err
}

// 减少video作者的获赞数
func (d *Dao) UserTotalFavoritedDown(VideoInfo model.VideoID) error {
	var video model.Video
	videoid := VideoInfo.VideoId
	err := d.Model(&model.Video{}).Where("id = ?", videoid).First(&video).Error
	if err != nil {
		return err
	}

	authorid := video.Author.Id
	err = d.Model(&model.User{}).Where("id = ?", authorid).UpdateColumn("total_favorited", gorm.Expr("total_favorited - ?", 1)).Error
	return err
}

// 减少用户的点赞数
func (d *Dao) UserFavoriteCountDown(UserInfo model.UserID) error {
	userid := UserInfo.UserId
	err := d.Model(&model.User{}).Where("id = ?", userid).UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error
	return err
}

// 查找LikeList的video
func (d *Dao) FindVideosInLikeList(UserInfo model.UserID) ([]model.Video, error) {
	//根据like_list表查找所有的video_id
	userid := UserInfo.UserId
	var lists []model.LikeList
	var ids []int64
	result := d.Model(&model.LikeList{}).Select("video_id").Where("user_id = ?", userid).Find(&lists)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, list := range lists {
		ids = append(ids, list.VideoId)
	}

	//由video_id返回所有的video
	var videos []model.Video
	out := d.Model(&model.Video{}).Where("id IN ?", ids).Find(&videos)
	if out.Error != nil {
		return nil, out.Error
	}
	return videos, out.Error
}
