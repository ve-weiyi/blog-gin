package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/logic"
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type PageHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewPageHandler(svcCtx *svctx.ServiceContext) *PageHandler {
	return &PageHandler{
		svcCtx: svcCtx,
	}
}

// 获取页面列表
// @Param data body types.QueryPageListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/page/query_page_list [POST]
func (h *PageHandler) QueryPageList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryPageListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewPageLogic(h.svcCtx)
	resp, err := l.QueryPageList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
