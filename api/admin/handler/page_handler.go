package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type PageHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewPageHandler(svcCtx *svctx.ServiceContext) *PageHandler {
	return &PageHandler{
		svcCtx: svcCtx,
	}
}

// 创建页面
// @Param data body types.CreatePageReq true "请求参数"
// @Success 200 {object} types.PageVO "返回信息"
// @Router /admin-api/v1/page/create_page [POST]
func (h *PageHandler) CreatePage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreatePageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewPageLogic(h.svcCtx)
	resp, err := l.CreatePage(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 删除页面
// @Param data body types.DeletePageReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/page/delete_page [DELETE]
func (h *PageHandler) DeletePage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeletePageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewPageLogic(h.svcCtx)
	resp, err := l.DeletePage(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取页面列表
// @Param data body types.QueryPageListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/page/query_page_list [POST]
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

// 更新页面
// @Param data body types.UpdatePageReq true "请求参数"
// @Success 200 {object} types.PageVO "返回信息"
// @Router /admin-api/v1/page/update_page [PUT]
func (h *PageHandler) UpdatePage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdatePageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewPageLogic(h.svcCtx)
	resp, err := l.UpdatePage(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
