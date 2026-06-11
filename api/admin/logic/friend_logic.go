package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
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

// 创建友链
func (l *FriendLogic) CreateFriend(reqCtx *request.Context, req *types.CreateFriendReq) (resp *types.FriendVO, err error) {
	// TODO: implement business logic

	return
}

// 批量删除友链
func (l *FriendLogic) DeleteFriend(reqCtx *request.Context, req *types.DeleteFriendReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取友链列表
func (l *FriendLogic) QueryFriendList(reqCtx *request.Context, req *types.QueryFriendListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 更新友链
func (l *FriendLogic) UpdateFriend(reqCtx *request.Context, req *types.UpdateFriendReq) (resp *types.FriendVO, err error) {
	// TODO: implement business logic

	return
}
