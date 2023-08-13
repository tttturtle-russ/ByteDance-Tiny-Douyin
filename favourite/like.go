package favourite

import (
	"ByteDance-Tiny-Douyin/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Like(userid int64, videoid int64, c *gin.Context) {
	//LikeList加入记录
	err := dao.RecordAdd(userid, videoid)
	if err != nil {
		log.Printf("添加新点赞id错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//video获赞数增加
	if err := dao.TotalAdd(videoid); err != nil {
		log.Printf("增加视频点赞总数错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//video作者获赞数增加
	if err := dao.AuthorAdd(videoid); err != nil {
		log.Printf("增加作者点赞总数错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//user点赞数增加
	if err := dao.UserAdd(userid); err != nil {
		log.Printf("用户增加点赞错误， %v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}
}
