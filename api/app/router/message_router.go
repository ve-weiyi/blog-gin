package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type MessageRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewMessageRouter(svcCtx *svctx.ServiceContext) *MessageRouter {
	return &MessageRouter{
		svcCtx: svcCtx,
	}
}

func (r *MessageRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/api/v1")

		h := handler.NewMessageHandler(r.svcCtx)
		// 获取留言列表
		g.POST("/message/query_message_list", h.QueryMessageList)
	}
	{
		g := rg.Group("/api/v1")
		g.Use(r.svcCtx.UserAuth)

		h := handler.NewMessageHandler(r.svcCtx)
		// 创建留言
		g.POST("/message/create_message", h.CreateMessage)
	}
}
