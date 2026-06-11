package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type RoleLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewRoleLogic(svcCtx *svctx.ServiceContext) *RoleLogic {
	return &RoleLogic{
		svcCtx: svcCtx,
	}
}

// 创建角色
func (l *RoleLogic) CreateRole(reqCtx *request.Context, req *types.CreateRoleReq) (resp *types.RoleVO, err error) {
	// TODO: implement business logic

	return
}

// 删除角色
func (l *RoleLogic) DeleteRole(reqCtx *request.Context, req *types.DeleteRoleReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取角色详情
func (l *RoleLogic) GetRole(reqCtx *request.Context, req *types.GetRoleReq) (resp *types.RoleVO, err error) {
	// TODO: implement business logic

	return
}

// 查询角色权限配置
func (l *RoleLogic) GetRolePermissions(reqCtx *request.Context, req *types.GetRolePermissionsReq) (resp *types.RolePermissionsResp, err error) {
	// TODO: implement business logic

	return
}

// 获取角色列表
func (l *RoleLogic) QueryRoleList(reqCtx *request.Context, req *types.QueryRoleListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 更新角色信息
func (l *RoleLogic) UpdateRole(reqCtx *request.Context, req *types.UpdateRoleReq) (resp *types.RoleVO, err error) {
	// TODO: implement business logic

	return
}

// 更新角色接口权限
func (l *RoleLogic) UpdateRoleApiPermissions(reqCtx *request.Context, req *types.UpdateRoleApiPermissionsReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 更新角色菜单权限
func (l *RoleLogic) UpdateRoleMenuPermissions(reqCtx *request.Context, req *types.UpdateRoleMenuPermissionsReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}
