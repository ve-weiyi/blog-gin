package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type RoleHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewRoleHandler(svcCtx *svctx.ServiceContext) *RoleHandler {
	return &RoleHandler{
		svcCtx: svcCtx,
	}
}

// 创建角色
// @Param data body types.CreateRoleReq true "请求参数"
// @Success 200 {object} types.RoleVO "返回信息"
// @Router /admin-api/v1/role/create_role [POST]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateRoleReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewRoleLogic(h.svcCtx)
	resp, err := l.CreateRole(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 删除角色
// @Param data body types.DeleteRoleReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/role/delete_role [DELETE]
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteRoleReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewRoleLogic(h.svcCtx)
	resp, err := l.DeleteRole(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取角色详情
// @Param data body types.GetRoleReq true "请求参数"
// @Success 200 {object} types.RoleVO "返回信息"
// @Router /admin-api/v1/role/get_role [GET]
func (h *RoleHandler) GetRole(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetRoleReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewRoleLogic(h.svcCtx)
	resp, err := l.GetRole(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 查询角色权限配置
// @Param data body types.GetRolePermissionsReq true "请求参数"
// @Success 200 {object} types.RolePermissionsResp "返回信息"
// @Router /admin-api/v1/role/get_role_permissions [GET]
func (h *RoleHandler) GetRolePermissions(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetRolePermissionsReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewRoleLogic(h.svcCtx)
	resp, err := l.GetRolePermissions(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取角色列表
// @Param data body types.QueryRoleListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/role/query_role_list [POST]
func (h *RoleHandler) QueryRoleList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryRoleListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewRoleLogic(h.svcCtx)
	resp, err := l.QueryRoleList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新角色信息
// @Param data body types.UpdateRoleReq true "请求参数"
// @Success 200 {object} types.RoleVO "返回信息"
// @Router /admin-api/v1/role/update_role [PUT]
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateRoleReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewRoleLogic(h.svcCtx)
	resp, err := l.UpdateRole(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新角色接口权限
// @Param data body types.UpdateRoleApiPermissionsReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/role/update_role_api_permissions [PUT]
func (h *RoleHandler) UpdateRoleApiPermissions(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateRoleApiPermissionsReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewRoleLogic(h.svcCtx)
	resp, err := l.UpdateRoleApiPermissions(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新角色菜单权限
// @Param data body types.UpdateRoleMenuPermissionsReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/role/update_role_menu_permissions [PUT]
func (h *RoleHandler) UpdateRoleMenuPermissions(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateRoleMenuPermissionsReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewRoleLogic(h.svcCtx)
	resp, err := l.UpdateRoleMenuPermissions(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
