package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
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
		g := rg.Group("/api/v1")

		h := handler.NewConfigHandler(r.svcCtx)
		// 获取博客前台首页信息
		g.GET("/blog", h.GetBlogHomeInfo)
		// 获取关于我的信息
		g.GET("/blog/get_about_me", h.GetAboutMe)
	}
}
