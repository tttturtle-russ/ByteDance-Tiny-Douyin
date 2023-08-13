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

const LikeYou int32 = 1
const DislikeYou int32 = 2
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
	userid := claim.Id
	videoid := req.VideoId

	//Like表中是否存在点赞记录
	userlike, err := dao.CheckLike(userid, videoid) //改1
	if err != nil {
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//新点赞情况
	if userlike == false && req.ActionType == LikeYou {
		Like(userid, videoid, c)
	}

	if userlike == true && req.ActionType == DislikeYou {
		DisLike(userid, videoid, c)
	}

	//is_favourite处理（不确定是不是这么处理）
	//if req.ActionType == LikeYou {
	//	video.IsFavorite = true
	//} else if req.ActionType == Unlike {
	//	video.IsFavorite = false
	//}

	c.JSON(http.StatusOK, ActionResponse{
		Code: ResponseOk,
		Msg:  "操作成功",
	})
}
