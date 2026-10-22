package logic

import (
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type CategoryLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewCategoryLogic(svcCtx *svctx.ServiceContext) *CategoryLogic {
	return &CategoryLogic{
		svcCtx: svcCtx,
	}
}

// 获取分类列表
func (l *CategoryLogic) QueryCategoryList(reqCtx *request.Context, req *types.QueryCategoryListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}
