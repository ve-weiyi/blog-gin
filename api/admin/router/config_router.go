package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type ConfigRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewConfigRouter(svcCtx *svctx.ServiceContext) *ConfigRouter {
	return &ConfigRouter{
		svcCtx: svcCtx,
	}
}

func (r *ConfigRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewConfigHandler(r.svcCtx)
		// 获取关于我的信息
		g.GET("/admin/get_about_me", h.GetAboutMe)
		// 获取服务器信息
		g.GET("/admin/get_system_state", h.GetSystemState)
		// 获取网站配置
		g.GET("/admin/get_website_config", h.GetWebsiteConfig)
		// 更新关于我的信息
		g.PUT("/admin/update_about_me", h.UpdateAboutMe)
		// 更新网站配置
		g.PUT("/admin/update_website_config", h.UpdateWebsiteConfig)
	}
}
