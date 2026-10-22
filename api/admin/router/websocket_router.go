package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type WebsocketRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewWebsocketRouter(svcCtx *svctx.ServiceContext) *WebsocketRouter {
	return &WebsocketRouter{
		svcCtx: svcCtx,
	}
}

func (r *WebsocketRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")

		h := handler.NewWebsocketHandler(r.svcCtx)
		// WebSocket消息
		g.GET("/websocket", h.Websocket)
	}
}
