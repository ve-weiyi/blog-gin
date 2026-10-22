package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
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
		g := rg.Group("/admin-api/v1")

		h := handler.NewGuestHandler(r.svcCtx)
		// 获取游客信息
		g.GET("/guest/get_guest", h.GetGuest)
	}
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewGuestHandler(r.svcCtx)
		// 获取访客列表
		g.POST("/guest/query_guest_list", h.QueryGuestList)
	}
}
