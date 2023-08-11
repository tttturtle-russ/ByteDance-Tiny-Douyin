package favourite

//完成点赞和取消赞的操作
import (
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

// 进行参数绑定req
func Action(c *gin.Context) {
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

	//token传到user里
	//检查videoId的是否在user中
	result := IsfavouriteByid(req.VideoId)
	if req.ActionType == Like {
		result.Update("is_favourite", Like)
	} else if req.ActionType == Unlike {
		result.Update("is_favourite", Unlike)
	}
}
