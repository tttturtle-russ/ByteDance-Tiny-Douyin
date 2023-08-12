package dao

import (
	"ByteDance-Tiny-Douyin/db"
	"ByteDance-Tiny-Douyin/models"
	"log"
)

func Check(userid int64, videoid int64) (models.Video, bool, error) {
	//返回的是video记录，userid是否存在like， err错误信息
	var video models.Video
	err := db.DB.Model(&models.Video{}).Where("video_id = ?", videoid).Find(&video)
	if err != nil {
		log.Printf("%v", err)
		return video, false, err.Error
	}

	//在video的userid中查询
	for _, id := range video.UserFavourites {
		if userid == id {
			return video, true, err.Error
		}
	}
	return video, false, err.Error
}
