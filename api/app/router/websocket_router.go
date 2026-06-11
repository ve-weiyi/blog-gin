package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
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
		g := rg.Group("/api/v1")

		h := handler.NewWebsocketHandler(r.svcCtx)
		// WebSocket消息
		g.GET("/websocket", h.Websocket)
	}
}
