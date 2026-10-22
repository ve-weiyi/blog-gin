package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type AuthRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewAuthRouter(svcCtx *svctx.ServiceContext) *AuthRouter {
	return &AuthRouter{
		svcCtx: svcCtx,
	}
}

func (r *AuthRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")

		h := handler.NewAuthHandler(r.svcCtx)
		// 获取验证码
		g.GET("/auth/captcha", h.GetCaptcha)
	}
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.RateLimit)

		h := handler.NewAuthHandler(r.svcCtx)
		// 邮箱验证码登录（仅登录）
		g.POST("/auth/login/email", h.EmailLogin)
		// 手机验证码登录（自动注册）
		g.POST("/auth/login/mobile", h.MobileLogin)
		// 密码登录（账号/手机号/邮箱）
		g.POST("/auth/login/password", h.PasswordLogin)
		// 获取第三方授权URL
		g.GET("/auth/oauth/:platform/authorize", h.GetOauthAuthorizeUrl)
		// 第三方登录（自动注册）
		g.POST("/auth/oauth/:platform/login", h.OauthLogin)
		// 重置密码
		g.POST("/auth/password/reset", h.ResetPassword)
		// 刷新token
		g.POST("/auth/refresh-token", h.RefreshToken)
		// 邮箱注册
		g.POST("/auth/register/email", h.EmailRegister)
		// 发送邮箱验证码
		g.POST("/auth/send/email", h.SendEmailCode)
		// 发送手机验证码
		g.POST("/auth/send/mobile", h.SendMobileCode)
	}
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)

		h := handler.NewAuthHandler(r.svcCtx)
		// 退出登录（需登录）
		g.POST("/auth/logout", h.Logout)
	}
}
