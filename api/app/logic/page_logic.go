package logic

import (
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type PageLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewPageLogic(svcCtx *svctx.ServiceContext) *PageLogic {
	return &PageLogic{
		svcCtx: svcCtx,
	}
}

// 获取页面列表
func (l *PageLogic) QueryPageList(reqCtx *request.Context, req *types.QueryPageListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}
