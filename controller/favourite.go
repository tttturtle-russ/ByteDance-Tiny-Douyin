package controller

//完成点赞和取消赞的操作
import (
	"ByteDance-Tiny-Douyin/model"
	"ByteDance-Tiny-Douyin/service"
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

type ListRequest struct {
	UserId int64
	Token  string
}

type ListResponse struct {
	Code      int32  // 状态码，0-成功，其他值-失败
	Msg       string // 返回状态描述
	VideoList []model.Video
}

const LikeYou int32 = 1
const DislikeYou int32 = 2
const ResponseOk = 0

func FavouriteAction(c *gin.Context) {
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
		log.Printf("解析错误，%v", err.Error())
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	userid := claim.Id
	videoid := req.VideoId

	//Like表中是否存在点赞记录
	svc := service.Newservice(c)
	userlike, err := svc.CheckLike(userid, videoid) //改1
	if err != nil {
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//新点赞情况
	if userlike == false && req.ActionType == LikeYou {
		LikeAction(userid, videoid, c)
	}

	if userlike == true && req.ActionType == DislikeYou {
		DislikeAction(userid, videoid, c)
	}

	//is_favourite处理（看需求）

	c.JSON(http.StatusOK, ActionResponse{
		Code: ResponseOk,
		Msg:  "操作成功",
	})
}

// 对列表的操作
func FavouriteList(c *gin.Context) {
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
	var like []model.Video
	svc := service.Newservice(c)
	like, err = svc.LikeList(userid)
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

func LikeAction(userid int64, videoid int64, c *gin.Context) {
	//LikeList加入记录
	svc := service.Newservice(c)
	err := svc.RecordAdd(userid, videoid)

	if err != nil {
		log.Printf("添加新点赞id错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//video获赞数增加
	svc_1 := service.Newservice(c)

	if err := svc_1.TotalAdd(videoid); err != nil {
		log.Printf("增加视频点赞总数错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//video作者获赞数增加
	svc_2 := service.Newservice(c)
	if err := svc_2.AuthorAdd(videoid); err != nil {
		log.Printf("增加作者点赞总数错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//user点赞数增加
	svc_3 := service.Newservice(c)
	if err := svc_3.UserAdd(userid); err != nil {
		log.Printf("用户增加点赞错误， %v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}
}

func DislikeAction(userid int64, videoid int64, c *gin.Context) {
	//删除点赞用户的id
	svc := service.Newservice(c)
	if err := svc.RecordDelete(userid, videoid); err != nil {
		log.Printf("删除点赞id错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//减少video总点赞数
	svc_1 := service.Newservice(c)
	if err := svc_1.TotalDown(videoid); err != nil {
		log.Printf("减少视频点赞总数错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//减少作者获赞数
	svc_2 := service.Newservice(c)
	if err := svc_2.AuthorDown(videoid); err != nil {
		log.Printf("减少作者点赞总数错误，%v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}

	//减少用户点赞数
	svc_3 := service.Newservice(c)
	if err := svc_3.UserDown(userid); err != nil {
		log.Printf("减少用户点赞错误， %v", err)
		res := &ActionResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, res)
	}
}
