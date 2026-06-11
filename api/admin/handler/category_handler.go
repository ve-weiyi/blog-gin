package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type CategoryHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewCategoryHandler(svcCtx *svctx.ServiceContext) *CategoryHandler {
	return &CategoryHandler{
		svcCtx: svcCtx,
	}
}

// 创建分类
// @Param data body types.CreateCategoryReq true "请求参数"
// @Success 200 {object} types.CategoryVO "返回信息"
// @Router /admin-api/v1/category/create_category [POST]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateCategoryReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewCategoryLogic(h.svcCtx)
	resp, err := l.CreateCategory(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 批量删除分类
// @Param data body types.DeleteCategoryReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/category/delete_category [DELETE]
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteCategoryReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewCategoryLogic(h.svcCtx)
	resp, err := l.DeleteCategory(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取分类列表
// @Param data body types.QueryCategoryListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/category/query_category_list [POST]
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

// 更新分类
// @Param data body types.UpdateCategoryReq true "请求参数"
// @Success 200 {object} types.CategoryVO "返回信息"
// @Router /admin-api/v1/category/update_category [PUT]
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateCategoryReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewCategoryLogic(h.svcCtx)
	resp, err := l.UpdateCategory(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
