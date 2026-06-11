package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/app/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type TagHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewTagHandler(svcCtx *svctx.ServiceContext) *TagHandler {
	return &TagHandler{
		svcCtx: svcCtx,
	}
}

// 获取标签列表
// @Param data body types.QueryTagListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/tag/query_tag_list [POST]
func (h *TagHandler) QueryTagList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryTagListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewTagLogic(h.svcCtx)
	resp, err := l.QueryTagList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
