package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/logic"
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type CategoryHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewCategoryHandler(svcCtx *svctx.ServiceContext) *CategoryHandler {
	return &CategoryHandler{
		svcCtx: svcCtx,
	}
}

// 获取分类列表
// @Param data body types.QueryCategoryListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/category/query_category_list [POST]
func (h *CategoryHandler) QueryCategoryList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryCategoryListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewCategoryLogic(h.svcCtx)
	resp, err := l.QueryCategoryList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
