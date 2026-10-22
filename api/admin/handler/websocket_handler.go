package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/logic"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type WebsocketHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewWebsocketHandler(svcCtx *svctx.ServiceContext) *WebsocketHandler {
	return &WebsocketHandler{
		svcCtx: svcCtx,
	}
}

// WebSocket消息
// @Router /admin-api/v1/websocket [GET]
func (h *WebsocketHandler) Websocket(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewWebsocketLogic(h.svcCtx)
	err = l.Websocket(reqCtx)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, nil)
}
