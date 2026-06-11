package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
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
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewMessageHandler(r.svcCtx)
		// 批量删除留言
		g.DELETE("/message/delete_message", h.DeleteMessage)
		// 获取留言列表
		g.POST("/message/query_message_list", h.QueryMessageList)
		// 批量更新留言状态
		g.PUT("/message/update_message_status", h.UpdateMessageStatus)
	}
}
