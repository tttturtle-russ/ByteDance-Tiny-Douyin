package favourite

import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/models"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ListRequest struct {
	UserId int64
	Token  string
}

type ListResponse struct {
	Code      int32
	Msg       string
	VideoList []models.Video
}

func List(c *gin.Context) {
	//参数绑定
	var req ListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("参数绑定错误，%v", err)
		res := &ListResponse{
			Code:      1,
			Msg:       err.Error(),
			VideoList: nil,
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//解析token
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		log.Printf("解析错误，%v", err)
		res := &ListResponse{
			Code:      1,
			Msg:       err.Error(),
			VideoList: nil,
		}
		c.JSON(http.StatusBadRequest, res)
	}

	userid := claim.Id
	if userid != req.UserId {
		log.Printf("ID匹配错误")
		res := &ListResponse{
			Code:      1,
			Msg:       "匹配错误",
			VideoList: nil,
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//查找喜欢列表
	var like []models.Video
	like, err = dao.LikeList(userid)
	if err != nil {
		log.Printf("查找错误，%v", err)
		res := &ListResponse{
			Code:      1,
			Msg:       "查找错误",
			VideoList: nil,
		}
		c.JSON(http.StatusBadRequest, res)
	}

	c.JSON(http.StatusOK, &ListResponse{
		Code:      0,
		Msg:       "返回成功",
		VideoList: like,
	})
}
