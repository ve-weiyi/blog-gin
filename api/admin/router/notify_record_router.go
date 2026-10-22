package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type NotifyRecordRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewNotifyRecordRouter(svcCtx *svctx.ServiceContext) *NotifyRecordRouter {
	return &NotifyRecordRouter{
		svcCtx: svcCtx,
	}
}

func (r *NotifyRecordRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewNotifyRecordHandler(r.svcCtx)
		// 批量标记投递记录为已读
		g.POST("/notify_record/batch_mark_records_read", h.BatchMarkRecordsRead)
		// 批量删除投递记录
		g.DELETE("/notify_record/delete_notify_record", h.DeleteNotifyRecord)
		// 全部标记投递记录为已读
		g.POST("/notify_record/mark_all_records_read", h.MarkAllRecordsRead)
		// 标记单条投递记录为已读
		g.POST("/notify_record/mark_record_read", h.MarkRecordRead)
		// 获取统一投递记录列表
		g.POST("/notify_record/query_notify_record_list", h.QueryNotifyRecordList)
		// 查询用户 inbox 投递列表
		g.POST("/notify_record/query_user_inbox_record_list", h.QueryUserInboxRecordList)
	}
}
