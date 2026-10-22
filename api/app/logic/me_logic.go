package logic

import (
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
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

// 停用账号（进入冷静期）
func (l *MeLogic) DeactivateAccount(reqCtx *request.Context, req *types.DeactivateAccountReq) (resp *types.DeactivateAccountResp, err error) {
	// TODO: implement business logic

	return
}

// 获取用户点赞集合
func (l *MeLogic) GetUserLike(reqCtx *request.Context, req *types.GetUserLikeReq) (resp *types.GetUserLikeResp, err error) {
	// TODO: implement business logic

	return
}

// 获取当前用户信息
func (l *MeLogic) GetUserProfile(reqCtx *request.Context, req *types.GetUserProfileReq) (resp *types.GetUserProfileResp, err error) {
	// TODO: implement business logic

	return
}

// 恢复账号（冷静期内）
func (l *MeLogic) ReactivateAccount(reqCtx *request.Context, req *types.ReactivateAccountReq) (resp *types.ReactivateAccountResp, err error) {
	// TODO: implement business logic

	return
}

// 解绑第三方平台账号
func (l *MeLogic) UnbindUserThirdParty(reqCtx *request.Context, req *types.UnbindUserThirdPartyReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 更新当前用户信息
func (l *MeLogic) UpdateUserProfile(reqCtx *request.Context, req *types.UpdateUserProfileReq) (resp *types.UpdateUserProfileResp, err error) {
	// TODO: implement business logic

	return
}
