package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/app/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type GuestHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewGuestHandler(svcCtx *svctx.ServiceContext) *GuestHandler {
	return &GuestHandler{
		svcCtx: svcCtx,
	}
}

// 获取游客信息
// @Param data body types.GetGuestReq true "请求参数"
// @Success 200 {object} types.GetGuestResp "返回信息"
// @Router /api/v1/guest/get_guest [GET]
func (h *GuestHandler) GetGuest(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetGuestReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewGuestLogic(h.svcCtx)
	resp, err := l.GetGuest(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
