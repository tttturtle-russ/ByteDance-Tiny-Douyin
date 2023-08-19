package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DisFollow(userid int64, disfollowid int64, c *gin.Context) {
	//删除点赞用户的id
	if err := dao.RecordDelete(userid, disfollowid); err != nil {
		log.Printf("取消关注错误，%v", err)
		res := &ControllerResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//减少粉丝数
	if err := dao.FollowerDown(disfollowid); err != nil {
		log.Printf("减少粉丝数错误，%v", err)
		res := &ControllerResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//减少关注数
	if err := dao.FollowDown(userid); err != nil {
		log.Printf("减少关注数错误，%v", err)
		res := &ControllerResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//移除关注和粉丝列表
	if err := dao.Down(userid); err != nil {
		log.Printf("移除列表错误， %v", err)
		res := &ControllerResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}
}
