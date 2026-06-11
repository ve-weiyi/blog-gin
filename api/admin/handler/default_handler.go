package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type DefaultHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewDefaultHandler(svcCtx *svctx.ServiceContext) *DefaultHandler {
	return &DefaultHandler{
		svcCtx: svcCtx,
	}
}

// ping
// @Param data body types.PingReq true "请求参数"
// @Success 200 {object} types.PingResp "返回信息"
// @Router /admin-api/v1/ping [GET]
func (h *DefaultHandler) Ping(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.PingReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewDefaultLogic(h.svcCtx)
	resp, err := l.Ping(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
