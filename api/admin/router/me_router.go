package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type MeRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewMeRouter(svcCtx *svctx.ServiceContext) *MeRouter {
	return &MeRouter{
		svcCtx: svcCtx,
	}
}

func (r *MeRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)

		h := handler.NewMeHandler(r.svcCtx)
		// 绑定邮箱
		g.POST("/user/me/bind_user_email", h.BindUserEmail)
		// 绑定手机号
		g.POST("/user/me/bind_user_phone", h.BindUserPhone)
		// 绑定第三方平台账号
		g.POST("/user/me/bind_user_third_party", h.BindUserThirdParty)
		// 获取用户接口权限
		g.GET("/user/me/get_user_apis", h.GetUserApis)
		// 获取用户菜单权限
		g.GET("/user/me/get_user_menus", h.GetUserMenus)
		// 获取当前用户信息
		g.GET("/user/me/get_user_profile", h.GetUserProfile)
		// 获取用户角色
		g.GET("/user/me/get_user_roles", h.GetUserRoles)
		// 查询用户登录历史
		g.POST("/user/me/query_user_login_history", h.QueryUserLoginHistory)
		// 解绑第三方平台账号
		g.POST("/user/me/unbind_user_third_party", h.UnbindUserThirdParty)
		// 修改用户头像
		g.PUT("/user/me/update_user_avatar", h.UpdateUserAvatar)
		// 修改用户密码
		g.PUT("/user/me/update_user_password", h.UpdateUserPassword)
		// 更新当前用户信息
		g.PUT("/user/me/update_user_profile", h.UpdateUserProfile)
	}
}
