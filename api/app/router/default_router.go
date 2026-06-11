package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
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
		g := rg.Group("/api/v1")

		h := handler.NewDefaultHandler(r.svcCtx)
		// 健康检查
		g.GET("/ping", h.Ping)
	}
}
