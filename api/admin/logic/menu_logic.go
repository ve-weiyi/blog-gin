package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type MenuLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewMenuLogic(svcCtx *svctx.ServiceContext) *MenuLogic {
	return &MenuLogic{
		svcCtx: svcCtx,
	}
}

// 清空菜单列表
func (l *MenuLogic) CleanMenu(reqCtx *request.Context, req *types.CleanMenuReq) (resp *types.CleanMenuResp, err error) {
	// TODO: implement business logic

	return
}

// 创建菜单
func (l *MenuLogic) CreateMenu(reqCtx *request.Context, req *types.CreateMenuReq) (resp *types.MenuVO, err error) {
	// TODO: implement business logic

	return
}

// 删除菜单
func (l *MenuLogic) DeleteMenu(reqCtx *request.Context, req *types.DeleteMenuReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取菜单详情
func (l *MenuLogic) GetMenu(reqCtx *request.Context, req *types.GetMenuReq) (resp *types.MenuVO, err error) {
	// TODO: implement business logic

	return
}

// 获取菜单列表
func (l *MenuLogic) QueryMenuList(reqCtx *request.Context, req *types.QueryMenuListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 同步菜单列表
func (l *MenuLogic) SyncMenu(reqCtx *request.Context, req *types.SyncMenuReq) (resp *types.SyncMenuResp, err error) {
	// TODO: implement business logic

	return
}

// 更新菜单
func (l *MenuLogic) UpdateMenu(reqCtx *request.Context, req *types.UpdateMenuReq) (resp *types.MenuVO, err error) {
	// TODO: implement business logic

	return
}
