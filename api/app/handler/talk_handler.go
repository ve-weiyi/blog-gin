package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/logic"
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type TalkHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewTalkHandler(svcCtx *svctx.ServiceContext) *TalkHandler {
	return &TalkHandler{
		svcCtx: svcCtx,
	}
}

// 获取说说详情
// @Param data body types.GetTalkReq true "请求参数"
// @Success 200 {object} types.Talk "返回信息"
// @Router /api/v1/talk/get_talk [GET]
func (h *TalkHandler) GetTalk(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetTalkReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewTalkLogic(h.svcCtx)
	resp, err := l.GetTalk(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取说说列表
// @Param data body types.QueryTalkListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/talk/query_talk_list [POST]
func (h *TalkHandler) QueryTalkList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryTalkListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewTalkLogic(h.svcCtx)
	resp, err := l.QueryTalkList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 点赞说说
// @Param data body types.LikeTalkReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /api/v1/talk/like_talk [POST]
func (h *TalkHandler) LikeTalk(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.LikeTalkReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewTalkLogic(h.svcCtx)
	resp, err := l.LikeTalk(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
