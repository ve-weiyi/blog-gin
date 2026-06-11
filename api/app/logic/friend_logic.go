package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/app/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type FriendLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewFriendLogic(svcCtx *svctx.ServiceContext) *FriendLogic {
	return &FriendLogic{
		svcCtx: svcCtx,
	}
}

// 获取友链列表
func (l *FriendLogic) QueryFriendList(reqCtx *request.Context, req *types.QueryFriendListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}
