package favourite

import (
	"ByteDance-Tiny-Douyin/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func DisLike(userid int64, videoid int64, c *gin.Context) {
	//删除点赞用户的id
	if err := dao.RecordDelete(userid, videoid); err != nil {
		log.Printf("删除点赞id错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//减少video总点赞数
	if err := dao.TotalDown(videoid); err != nil {
		log.Printf("减少视频点赞总数错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//减少作者获赞数
	if err := dao.AuthorDown(videoid); err != nil {
		log.Printf("减少作者点赞总数错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//减少用户点赞数
	if err := dao.UserDown(userid); err != nil {
		log.Printf("减少用户点赞错误， %v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}
}
