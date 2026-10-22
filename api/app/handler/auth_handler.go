package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/logic"
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type AuthHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewAuthHandler(svcCtx *svctx.ServiceContext) *AuthHandler {
	return &AuthHandler{
		svcCtx: svcCtx,
	}
}

// 获取验证码
// @Param data body types.GetCaptchaReq true "请求参数"
// @Success 200 {object} types.GetCaptchaResp "返回信息"
// @Router /api/v1/auth/captcha [GET]
func (h *AuthHandler) GetCaptcha(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetCaptchaReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.GetCaptcha(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 邮箱验证码登录（仅登录）
// @Param data body types.EmailLoginReq true "请求参数"
// @Success 200 {object} types.LoginResp "返回信息"
// @Router /api/v1/auth/login/email [POST]
func (h *AuthHandler) EmailLogin(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.EmailLoginReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.EmailLogin(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 手机验证码登录（自动注册）
// @Param data body types.MobileLoginReq true "请求参数"
// @Success 200 {object} types.LoginResp "返回信息"
// @Router /api/v1/auth/login/mobile [POST]
func (h *AuthHandler) MobileLogin(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.MobileLoginReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.MobileLogin(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 密码登录（账号/手机号/邮箱）
// @Param data body types.PasswordLoginReq true "请求参数"
// @Success 200 {object} types.LoginResp "返回信息"
// @Router /api/v1/auth/login/password [POST]
func (h *AuthHandler) PasswordLogin(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.PasswordLoginReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.PasswordLogin(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取第三方授权URL
// @Param data body types.GetOauthAuthorizeUrlReq true "请求参数"
// @Success 200 {object} types.GetOauthAuthorizeUrlResp "返回信息"
// @Router /api/v1/auth/oauth/:platform/authorize [GET]
func (h *AuthHandler) GetOauthAuthorizeUrl(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetOauthAuthorizeUrlReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.GetOauthAuthorizeUrl(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 第三方登录（自动注册）
// @Param data body types.OauthLoginReq true "请求参数"
// @Success 200 {object} types.LoginResp "返回信息"
// @Router /api/v1/auth/oauth/:platform/login [POST]
func (h *AuthHandler) OauthLogin(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.OauthLoginReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.OauthLogin(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 重置密码
// @Param data body types.ResetPasswordReq true "请求参数"
// @Success 200 {object} types.ResetPasswordResp "返回信息"
// @Router /api/v1/auth/password/reset [POST]
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.ResetPasswordReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.ResetPassword(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 刷新token
// @Param data body types.RefreshTokenReq true "请求参数"
// @Success 200 {object} types.LoginResp "返回信息"
// @Router /api/v1/auth/refresh-token [POST]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.RefreshTokenReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.RefreshToken(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 邮箱注册
// @Param data body types.EmailRegisterReq true "请求参数"
// @Success 200 {object} types.EmailRegisterResp "返回信息"
// @Router /api/v1/auth/register/email [POST]
func (h *AuthHandler) EmailRegister(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.EmailRegisterReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.EmailRegister(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 发送邮箱验证码
// @Param data body types.SendEmailCodeReq true "请求参数"
// @Success 200 {object} types.SendEmailCodeResp "返回信息"
// @Router /api/v1/auth/send/email [POST]
func (h *AuthHandler) SendEmailCode(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.SendEmailCodeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.SendEmailCode(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 发送手机验证码
// @Param data body types.SendMobileCodeReq true "请求参数"
// @Success 200 {object} types.SendMobileCodeResp "返回信息"
// @Router /api/v1/auth/send/mobile [POST]
func (h *AuthHandler) SendMobileCode(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.SendMobileCodeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.SendMobileCode(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 退出登录（需登录）
// @Param data body types.LogoutReq true "请求参数"
// @Success 200 {object} types.LogoutResp "返回信息"
// @Router /api/v1/auth/logout [POST]
func (h *AuthHandler) Logout(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.LogoutReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAuthLogic(h.svcCtx)
	resp, err := l.Logout(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
