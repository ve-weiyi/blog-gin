package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type PageLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewPageLogic(svcCtx *svctx.ServiceContext) *PageLogic {
	return &PageLogic{
		svcCtx: svcCtx,
	}
}

// 创建页面
func (l *PageLogic) CreatePage(reqCtx *request.Context, req *types.CreatePageReq) (resp *types.PageVO, err error) {
	// TODO: implement business logic

	return
}

// 删除页面
func (l *PageLogic) DeletePage(reqCtx *request.Context, req *types.DeletePageReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取页面列表
func (l *PageLogic) QueryPageList(reqCtx *request.Context, req *types.QueryPageListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 更新页面
func (l *PageLogic) UpdatePage(reqCtx *request.Context, req *types.UpdatePageReq) (resp *types.PageVO, err error) {
	// TODO: implement business logic

	return
}
