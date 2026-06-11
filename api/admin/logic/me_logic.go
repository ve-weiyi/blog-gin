package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type MeLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewMeLogic(svcCtx *svctx.ServiceContext) *MeLogic {
	return &MeLogic{
		svcCtx: svcCtx,
	}
}

// 绑定邮箱
func (l *MeLogic) BindUserEmail(reqCtx *request.Context, req *types.BindUserEmailReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 绑定手机号
func (l *MeLogic) BindUserPhone(reqCtx *request.Context, req *types.BindUserPhoneReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 绑定第三方平台账号
func (l *MeLogic) BindUserThirdParty(reqCtx *request.Context, req *types.BindUserThirdPartyReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 获取用户接口权限
func (l *MeLogic) GetUserApis(reqCtx *request.Context, req *types.EmptyReq) (resp *types.GetUserApisResp, err error) {
	// TODO: implement business logic

	return
}

// 获取用户菜单权限
func (l *MeLogic) GetUserMenus(reqCtx *request.Context, req *types.EmptyReq) (resp *types.GetUserMenusResp, err error) {
	// TODO: implement business logic

	return
}

// 获取当前用户信息
func (l *MeLogic) GetUserProfile(reqCtx *request.Context, req *types.GetUserProfileReq) (resp *types.UserProfile, err error) {
	// TODO: implement business logic

	return
}

// 获取用户角色
func (l *MeLogic) GetUserRoles(reqCtx *request.Context, req *types.EmptyReq) (resp *types.GetUserRolesResp, err error) {
	// TODO: implement business logic

	return
}

// 查询用户登录历史
func (l *MeLogic) QueryUserLoginHistory(reqCtx *request.Context, req *types.QueryUserLoginHistoryReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 解绑第三方平台账号
func (l *MeLogic) UnbindUserThirdParty(reqCtx *request.Context, req *types.UnbindUserThirdPartyReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 修改用户头像
func (l *MeLogic) UpdateUserAvatar(reqCtx *request.Context, req *types.UpdateUserAvatarReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 修改用户密码
func (l *MeLogic) UpdateUserPassword(reqCtx *request.Context, req *types.UpdateUserPasswordReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 更新当前用户信息
func (l *MeLogic) UpdateUserProfile(reqCtx *request.Context, req *types.UpdateUserProfileReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}
