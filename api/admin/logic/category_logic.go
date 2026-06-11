package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type CategoryLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewCategoryLogic(svcCtx *svctx.ServiceContext) *CategoryLogic {
	return &CategoryLogic{
		svcCtx: svcCtx,
	}
}

// 创建分类
func (l *CategoryLogic) CreateCategory(reqCtx *request.Context, req *types.CreateCategoryReq) (resp *types.CategoryVO, err error) {
	// TODO: implement business logic

	return
}

// 批量删除分类
func (l *CategoryLogic) DeleteCategory(reqCtx *request.Context, req *types.DeleteCategoryReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取分类列表
func (l *CategoryLogic) QueryCategoryList(reqCtx *request.Context, req *types.QueryCategoryListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 更新分类
func (l *CategoryLogic) UpdateCategory(reqCtx *request.Context, req *types.UpdateCategoryReq) (resp *types.CategoryVO, err error) {
	// TODO: implement business logic

	return
}
