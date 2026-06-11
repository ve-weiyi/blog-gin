package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/app/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type AuthLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewAuthLogic(svcCtx *svctx.ServiceContext) *AuthLogic {
	return &AuthLogic{
		svcCtx: svcCtx,
	}
}

// 获取验证码
func (l *AuthLogic) GetCaptcha(reqCtx *request.Context, req *types.GetCaptchaReq) (resp *types.GetCaptchaResp, err error) {
	// TODO: implement business logic

	return
}

// 邮箱验证码登录（仅登录）
func (l *AuthLogic) EmailLogin(reqCtx *request.Context, req *types.EmailLoginReq) (resp *types.LoginResp, err error) {
	// TODO: implement business logic

	return
}

// 手机验证码登录（自动注册）
func (l *AuthLogic) MobileLogin(reqCtx *request.Context, req *types.MobileLoginReq) (resp *types.LoginResp, err error) {
	// TODO: implement business logic

	return
}

// 密码登录（账号/手机号/邮箱）
func (l *AuthLogic) PasswordLogin(reqCtx *request.Context, req *types.PasswordLoginReq) (resp *types.LoginResp, err error) {
	// TODO: implement business logic

	return
}

// 获取第三方授权URL
func (l *AuthLogic) GetOauthAuthorizeUrl(reqCtx *request.Context, req *types.GetOauthAuthorizeUrlReq) (resp *types.GetOauthAuthorizeUrlResp, err error) {
	// TODO: implement business logic

	return
}

// 第三方登录（自动注册）
func (l *AuthLogic) OauthLogin(reqCtx *request.Context, req *types.OauthLoginReq) (resp *types.LoginResp, err error) {
	// TODO: implement business logic

	return
}

// 重置密码
func (l *AuthLogic) ResetPassword(reqCtx *request.Context, req *types.ResetPasswordReq) (resp *types.ResetPasswordResp, err error) {
	// TODO: implement business logic

	return
}

// 刷新token
func (l *AuthLogic) RefreshToken(reqCtx *request.Context, req *types.RefreshTokenReq) (resp *types.LoginResp, err error) {
	// TODO: implement business logic

	return
}

// 邮箱注册
func (l *AuthLogic) EmailRegister(reqCtx *request.Context, req *types.EmailRegisterReq) (resp *types.EmailRegisterResp, err error) {
	// TODO: implement business logic

	return
}

// 发送邮箱验证码
func (l *AuthLogic) SendEmailCode(reqCtx *request.Context, req *types.SendEmailCodeReq) (resp *types.SendEmailCodeResp, err error) {
	// TODO: implement business logic

	return
}

// 发送手机验证码
func (l *AuthLogic) SendMobileCode(reqCtx *request.Context, req *types.SendMobileCodeReq) (resp *types.SendMobileCodeResp, err error) {
	// TODO: implement business logic

	return
}

// 退出登录（需登录）
func (l *AuthLogic) Logout(reqCtx *request.Context, req *types.LogoutReq) (resp *types.LogoutResp, err error) {
	// TODO: implement business logic

	return
}
