package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/handler"
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
		g := rg.Group("/api/v1")
		g.Use(r.svcCtx.UserAuth)

		h := handler.NewMeHandler(r.svcCtx)
		// 绑定邮箱
		g.POST("/user/me/bind_user_email", h.BindUserEmail)
		// 绑定手机号
		g.POST("/user/me/bind_user_phone", h.BindUserPhone)
		// 绑定第三方平台账号
		g.POST("/user/me/bind_user_third_party", h.BindUserThirdParty)
		// 停用账号（进入冷静期）
		g.POST("/user/me/deactivate_account", h.DeactivateAccount)
		// 获取用户点赞集合
		g.GET("/user/me/get_user_like", h.GetUserLike)
		// 获取当前用户信息
		g.GET("/user/me/get_user_profile", h.GetUserProfile)
		// 恢复账号（冷静期内）
		g.POST("/user/me/reactivate_account", h.ReactivateAccount)
		// 解绑第三方平台账号
		g.POST("/user/me/unbind_user_third_party", h.UnbindUserThirdParty)
		// 更新当前用户信息
		g.PUT("/user/me/update_user_profile", h.UpdateUserProfile)
	}
}
