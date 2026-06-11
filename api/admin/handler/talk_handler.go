package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type TalkHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewTalkHandler(svcCtx *svctx.ServiceContext) *TalkHandler {
	return &TalkHandler{
		svcCtx: svcCtx,
	}
}

// 创建说说
// @Param data body types.CreateTalkReq true "请求参数"
// @Success 200 {object} types.TalkVO "返回信息"
// @Router /admin-api/v1/talk/create_talk [POST]
func (h *TalkHandler) CreateTalk(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateTalkReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewTalkLogic(h.svcCtx)
	resp, err := l.CreateTalk(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 删除说说
// @Param data body types.DeleteTalkReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/talk/delete_talk [DELETE]
func (h *TalkHandler) DeleteTalk(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteTalkReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewTalkLogic(h.svcCtx)
	resp, err := l.DeleteTalk(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取说说详情
// @Param data body types.GetTalkReq true "请求参数"
// @Success 200 {object} types.TalkVO "返回信息"
// @Router /admin-api/v1/talk/get_talk [GET]
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
// @Router /admin-api/v1/talk/query_talk_list [POST]
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

// 更新说说
// @Param data body types.UpdateTalkReq true "请求参数"
// @Success 200 {object} types.TalkVO "返回信息"
// @Router /admin-api/v1/talk/update_talk [PUT]
func (h *TalkHandler) UpdateTalk(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateTalkReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewTalkLogic(h.svcCtx)
	resp, err := l.UpdateTalk(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
