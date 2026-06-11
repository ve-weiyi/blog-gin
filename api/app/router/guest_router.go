package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type GuestRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewGuestRouter(svcCtx *svctx.ServiceContext) *GuestRouter {
	return &GuestRouter{
		svcCtx: svcCtx,
	}
}

func (r *GuestRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/api/v1")

		h := handler.NewGuestHandler(r.svcCtx)
		// 获取游客信息
		g.GET("/guest/get_guest", h.GetGuest)
	}
}
