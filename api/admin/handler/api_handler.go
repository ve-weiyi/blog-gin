package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/logic"
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type ApiHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewApiHandler(svcCtx *svctx.ServiceContext) *ApiHandler {
	return &ApiHandler{
		svcCtx: svcCtx,
	}
}

// 清空接口列表
// @Param data body types.EmptyReq true "请求参数"
// @Success 200 {object} types.CleanApiResp "返回信息"
// @Router /admin-api/v1/api/clean_api [POST]
func (h *ApiHandler) CleanApi(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewApiLogic(h.svcCtx)
	resp, err := l.CleanApi(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 创建接口
// @Param data body types.CreateApiReq true "请求参数"
// @Success 200 {object} types.ApiVO "返回信息"
// @Router /admin-api/v1/api/create_api [POST]
func (h *ApiHandler) CreateApi(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateApiReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewApiLogic(h.svcCtx)
	resp, err := l.CreateApi(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 删除接口
// @Param data body types.DeleteApiReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/api/delete_api [DELETE]
func (h *ApiHandler) DeleteApi(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteApiReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewApiLogic(h.svcCtx)
	resp, err := l.DeleteApi(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取接口详情
// @Param data body types.GetApiReq true "请求参数"
// @Success 200 {object} types.ApiVO "返回信息"
// @Router /admin-api/v1/api/get_api [GET]
func (h *ApiHandler) GetApi(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetApiReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewApiLogic(h.svcCtx)
	resp, err := l.GetApi(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取接口列表
// @Param data body types.QueryApiListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/api/query_api_list [POST]
func (h *ApiHandler) QueryApiList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryApiListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewApiLogic(h.svcCtx)
	resp, err := l.QueryApiList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 同步接口列表
// @Param data body types.EmptyReq true "请求参数"
// @Success 200 {object} types.SyncApiResp "返回信息"
// @Router /admin-api/v1/api/sync_api [POST]
func (h *ApiHandler) SyncApi(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewApiLogic(h.svcCtx)
	resp, err := l.SyncApi(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新接口
// @Param data body types.UpdateApiReq true "请求参数"
// @Success 200 {object} types.ApiVO "返回信息"
// @Router /admin-api/v1/api/update_api [PUT]
func (h *ApiHandler) UpdateApi(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateApiReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewApiLogic(h.svcCtx)
	resp, err := l.UpdateApi(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
