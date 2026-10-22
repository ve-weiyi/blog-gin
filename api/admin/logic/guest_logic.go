package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
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

// 获取访客列表
func (l *GuestLogic) QueryGuestList(reqCtx *request.Context, req *types.QueryGuestListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}
