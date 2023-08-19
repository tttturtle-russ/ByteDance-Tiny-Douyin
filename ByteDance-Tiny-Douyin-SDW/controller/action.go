package controller

//完成关注和取消关注的操作
import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ControllerRequest struct {
	Token      string `json:"token"`
	UserId    int64  `form:"user_id"`
	ActionType int32  `from:"action_type"`
}

type ControllerResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

const FollowYou int32 = 1
const DisFollowYou int32 = 2
const ResponseOk = 0

func Controller(c *gin.Context) {
	// 进行参数绑定req
	var req ControllerRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("绑定失败")
		res := &ControllerResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}
	//解析token得userid
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		log.Printf("解析错误，%v", err.Error())
		res := &ControllerResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	userid := claim.Id
	followid := req.UserId
	
	//follow表中是否存在关注记录
	userfollow, err := dao.CheckFollow(userid, followid) 
	if err != nil {
		res := &ControllerResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}
	if userfollow == false && req.ActionType == FollowYou {
		Follow(userid, followid, c)
	}

	if userfollow == true && req.ActionType == DisFollowYou {
		DisFollow(userid, followid, c)
	}
