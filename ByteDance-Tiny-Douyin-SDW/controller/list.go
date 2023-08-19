package controller

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/models"
	"ByteDance-Tiny-Douyin/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListRequest struct {
	UserId int64
	Token  string
}

type ListResponse struct {
	Code         int32  // 状态码，0-成功，其他值-失败
	Msg          string // 返回状态描述
	FollowList   []models.User
	FollowerList []models.User
}

func List(c *gin.Context) {
	//参数绑定
	var req ListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("参数绑定错误，%v", err)
		res := &ListResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//解析token
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		log.Printf("解析错误，%v", err)
		res := &ListResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	userid := claim.Id
	if userid != req.UserId {
		log.Printf("ID匹配错误")
		res := &ListResponse{
			Code: 1,
			Msg:  "匹配错误",
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//查找关注列表
	var follow []models.User
	follow, err = dao.FollowList(userid)
	if err != nil {
		log.Printf("查找错误，%v", err)
		res := &ListResponse{
			Code: 1,
			Msg:  "查找错误",
		}
		c.JSON(http.StatusBadRequest, res)
	}

	c.JSON(http.StatusOK, &ListResponse{
		Code:       0,
		Msg:        "返回成功",
		FollowList: follow,
	})
	//查找粉丝列表
	var follower []models.User
	follower, err = dao.FollowList(userid)
	if err != nil {
		log.Printf("查找错误，%v", err)
		res := &ListResponse{
			Code: 1,
			Msg:  "查找错误",
		}
		c.JSON(http.StatusBadRequest, res)
	}

	c.JSON(http.StatusOK, &ListResponse{
		Code:         0,
		Msg:          "返回成功",
		FollowerList: follower,
	})
}
