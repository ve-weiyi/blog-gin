package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type NotifyMessageRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewNotifyMessageRouter(svcCtx *svctx.ServiceContext) *NotifyMessageRouter {
	return &NotifyMessageRouter{
		svcCtx: svcCtx,
	}
}

func (r *NotifyMessageRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewNotifyMessageHandler(r.svcCtx)
		// 创建统一通知消息
		g.POST("/notify_message/create_notify_message", h.CreateNotifyMessage)
		// 批量删除统一通知消息
		g.DELETE("/notify_message/delete_notify_message", h.DeleteNotifyMessage)
		// 获取统一通知消息详情
		g.GET("/notify_message/get_notify_message", h.GetNotifyMessage)
		// 发布统一通知消息
		g.POST("/notify_message/publish_notify_message", h.PublishNotifyMessage)
		// 获取统一通知消息列表
		g.POST("/notify_message/query_notify_message_list", h.QueryNotifyMessageList)
		// 撤回统一通知消息
		g.POST("/notify_message/revoke_notify_message", h.RevokeNotifyMessage)
		// 更新统一通知消息
		g.PUT("/notify_message/update_notify_message", h.UpdateNotifyMessage)
	}
}
