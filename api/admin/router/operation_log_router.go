package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type OperationLogRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewOperationLogRouter(svcCtx *svctx.ServiceContext) *OperationLogRouter {
	return &OperationLogRouter{
		svcCtx: svcCtx,
	}
}

func (r *OperationLogRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewOperationLogHandler(r.svcCtx)
		// 批量删除操作日志
		g.DELETE("/operation_log/delete_operation_log", h.DeleteOperationLog)
		// 获取操作日志列表
		g.POST("/operation_log/query_operation_log_list", h.QueryOperationLogList)
	}
}
