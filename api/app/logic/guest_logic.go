package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/app/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type GuestLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewGuestLogic(svcCtx *svctx.ServiceContext) *GuestLogic {
	return &GuestLogic{
		svcCtx: svcCtx,
	}
}

// 获取游客信息
func (l *GuestLogic) GetGuest(reqCtx *request.Context, req *types.GetGuestReq) (resp *types.GetGuestResp, err error) {
	// TODO: implement business logic

	return
}
