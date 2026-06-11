package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type UserLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewUserLogic(svcCtx *svctx.ServiceContext) *UserLogic {
	return &UserLogic{
		svcCtx: svcCtx,
	}
}

// 获取在线用户列表
func (l *UserLogic) GetOnlineUsers(reqCtx *request.Context, req *types.EmptyReq) (resp *types.OnlineUserListResp, err error) {
	// TODO: implement business logic

	return
}

// 获取用户详情
func (l *UserLogic) GetUserDetail(reqCtx *request.Context, req *types.GetUserDetailReq) (resp *types.GetUserDetailResp, err error) {
	// TODO: implement business logic

	return
}

// 获取用户列表
func (l *UserLogic) QueryUserList(reqCtx *request.Context, req *types.QueryUserListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 重置用户密码
func (l *UserLogic) ResetUserPassword(reqCtx *request.Context, req *types.ResetUserPasswordReq) (resp *types.ResetUserPasswordResp, err error) {
	// TODO: implement business logic

	return
}

// 更新用户角色
func (l *UserLogic) UpdateUserRoles(reqCtx *request.Context, req *types.UpdateUserRolesReq) (resp *types.UpdateUserRolesResp, err error) {
	// TODO: implement business logic

	return
}

// 更新用户状态
func (l *UserLogic) UpdateUserStatus(reqCtx *request.Context, req *types.UpdateUserStatusReq) (resp *types.UpdateUserStatusResp, err error) {
	// TODO: implement business logic

	return
}
