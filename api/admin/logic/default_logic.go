package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type DefaultLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewDefaultLogic(svcCtx *svctx.ServiceContext) *DefaultLogic {
	return &DefaultLogic{
		svcCtx: svcCtx,
	}
}

// ping
func (l *DefaultLogic) Ping(reqCtx *request.Context, req *types.PingReq) (resp *types.PingResp, err error) {
	// TODO: implement business logic

	return
}
