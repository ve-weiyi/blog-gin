package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/app/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type TagLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewTagLogic(svcCtx *svctx.ServiceContext) *TagLogic {
	return &TagLogic{
		svcCtx: svcCtx,
	}
}

// 获取标签列表
func (l *TagLogic) QueryTagList(reqCtx *request.Context, req *types.QueryTagListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}
