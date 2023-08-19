package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Follow(userid int64, followid int64, c *gin.Context) {
	//FollowList加入记录
	err := dao.RecordAdd(userid, followid)
	if err != nil {
		log.Printf("关注失败，%v", err)
		res := &ControllerResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//粉丝数增加
	if err := dao.FollowerAdd(followid); err != nil {
		log.Printf("增加粉丝总数错误，%v", err)
		res := &ControllerResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//关注数增加
	if err := dao.FollowAdd(userid); err != nil {
		log.Printf("增加关注总数错误，%v", err)
		res := &ControllerResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//加入粉丝列表和关注列表
	if err := dao.Add(userid, followid); err != nil {
		log.Printf("关注错误， %v", err)
		res := &ControllerResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}
}
