package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/logic"
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type MenuHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewMenuHandler(svcCtx *svctx.ServiceContext) *MenuHandler {
	return &MenuHandler{
		svcCtx: svcCtx,
	}
}

// 清空菜单列表
// @Param data body types.CleanMenuReq true "请求参数"
// @Success 200 {object} types.CleanMenuResp "返回信息"
// @Router /admin-api/v1/menu/clean_menu [POST]
func (h *MenuHandler) CleanMenu(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CleanMenuReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMenuLogic(h.svcCtx)
	resp, err := l.CleanMenu(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 创建菜单
// @Param data body types.CreateMenuReq true "请求参数"
// @Success 200 {object} types.MenuVO "返回信息"
// @Router /admin-api/v1/menu/create_menu [POST]
func (h *MenuHandler) CreateMenu(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateMenuReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMenuLogic(h.svcCtx)
	resp, err := l.CreateMenu(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 删除菜单
// @Param data body types.DeleteMenuReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/menu/delete_menu [DELETE]
func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteMenuReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMenuLogic(h.svcCtx)
	resp, err := l.DeleteMenu(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取菜单详情
// @Param data body types.GetMenuReq true "请求参数"
// @Success 200 {object} types.MenuVO "返回信息"
// @Router /admin-api/v1/menu/get_menu [GET]
func (h *MenuHandler) GetMenu(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetMenuReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMenuLogic(h.svcCtx)
	resp, err := l.GetMenu(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取菜单列表
// @Param data body types.QueryMenuListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/menu/query_menu_list [POST]
func (h *MenuHandler) QueryMenuList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryMenuListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMenuLogic(h.svcCtx)
	resp, err := l.QueryMenuList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 同步菜单列表
// @Param data body types.SyncMenuReq true "请求参数"
// @Success 200 {object} types.SyncMenuResp "返回信息"
// @Router /admin-api/v1/menu/sync_menu [POST]
func (h *MenuHandler) SyncMenu(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.SyncMenuReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMenuLogic(h.svcCtx)
	resp, err := l.SyncMenu(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新菜单
// @Param data body types.UpdateMenuReq true "请求参数"
// @Success 200 {object} types.MenuVO "返回信息"
// @Router /admin-api/v1/menu/update_menu [PUT]
func (h *MenuHandler) UpdateMenu(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateMenuReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMenuLogic(h.svcCtx)
	resp, err := l.UpdateMenu(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
