package favourite

//完成点赞和取消赞的操作
import (
	"ByteDance-Tiny-Douyin/dao"
	"ByteDance-Tiny-Douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ActionRequest struct {
	Token      string `json:"token"`
	VideoId    int64  `form:"video_id"`
	ActionType int32  `from:"action_type"`
}

type ActionResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

const Like int32 = 1
const Unlike int32 = 2
const ResponseOk = 0

func Action(c *gin.Context) {
	// 进行参数绑定req
	var req ActionRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Printf("绑定失败")
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//解析token得userid
	claim, err := util.ParseToken(req.Token)
	if err != nil {
		log.Printf("%v", err.Error())
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}
	userid := claim.Id     //复制用户id
	videoid := req.VideoId //复制videoid
	video, userlike, err := dao.Check(userid, videoid)
	if err != nil {
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//新点赞情况
	if userlike == false && req.ActionType == Like {
		//video的点赞id增加一条
		if err := dao.Insert(userid, videoid); err != nil {
			log.Printf("添加新点赞id错误，%v", err)
			res := &ActionResponse{
				Code: 1,
				Msg:  err.Error(),
			}
			c.JSON(http.StatusBadRequest, res)
		}

		//video点赞数增加
		if err := dao.TotalAdd(videoid); err != nil {
			log.Printf("增加视频点赞总数错误，%v", err)
			res := &ActionResponse{
				Code: 1,
				Msg:  err.Error(),
			}
			c.JSON(http.StatusBadRequest, res)
		}

		//video的作者总获赞数增加
		if err := dao.AuthorAdd(videoid); err != nil {
			log.Printf("增加作者点赞总数错误，%v", err)
			res := &ActionResponse{
				Code: 1,
				Msg:  err.Error(),
			}
			c.JSON(http.StatusBadRequest, res)
		}

		//点赞用户点赞数增加
		if err := dao.UserAdd(userid); err != nil {
			log.Printf("用户增加点赞错误， %v", err)
			res := &ActionResponse{
				Code: 1,
				Msg:  err.Error(),
			}
			c.JSON(http.StatusBadRequest, res)
		}
	}

	if userlike == true && req.ActionType == Unlike {
		//删除取消点赞的id
		if err := dao.Delete(userid, videoid); err != nil {
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

	//is_favourite处理（不确定是不是这么处理）
	if req.ActionType == Like {
		video.IsFavorite = true
	} else if req.ActionType == Unlike {
		video.IsFavorite = false
	}

	c.JSON(http.StatusOK, ActionResponse{
		Code: ResponseOk,
		Msg:  "操作成功",
	})
}
