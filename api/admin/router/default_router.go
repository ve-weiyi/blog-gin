package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type DefaultRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewDefaultRouter(svcCtx *svctx.ServiceContext) *DefaultRouter {
	return &DefaultRouter{
		svcCtx: svcCtx,
	}
}

func (r *DefaultRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")

		h := handler.NewDefaultHandler(r.svcCtx)
		// ping
		g.GET("/ping", h.Ping)
	}
}
