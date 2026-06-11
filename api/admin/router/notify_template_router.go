package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type NotifyTemplateRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewNotifyTemplateRouter(svcCtx *svctx.ServiceContext) *NotifyTemplateRouter {
	return &NotifyTemplateRouter{
		svcCtx: svcCtx,
	}
}

func (r *NotifyTemplateRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewNotifyTemplateHandler(r.svcCtx)
		// 创建通知模板
		g.POST("/notify_template/create_notify_template", h.CreateNotifyTemplate)
		// 批量删除通知模板
		g.DELETE("/notify_template/delete_notify_template", h.DeleteNotifyTemplate)
		// 获取通知模板详情
		g.GET("/notify_template/get_notify_template", h.GetNotifyTemplate)
		// 获取通知模板列表
		g.POST("/notify_template/query_notify_template_list", h.QueryNotifyTemplateList)
		// 更新通知模板
		g.PUT("/notify_template/update_notify_template", h.UpdateNotifyTemplate)
	}
}
