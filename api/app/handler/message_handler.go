package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/logic"
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type MessageHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewMessageHandler(svcCtx *svctx.ServiceContext) *MessageHandler {
	return &MessageHandler{
		svcCtx: svcCtx,
	}
}

// 获取留言列表
// @Param data body types.QueryMessageListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/message/query_message_list [POST]
func (h *MessageHandler) QueryMessageList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryMessageListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMessageLogic(h.svcCtx)
	resp, err := l.QueryMessageList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 创建留言
// @Param data body types.CreateMessageReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /api/v1/message/create_message [POST]
func (h *MessageHandler) CreateMessage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateMessageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMessageLogic(h.svcCtx)
	resp, err := l.CreateMessage(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
