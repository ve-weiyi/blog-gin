package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type VisitLogRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewVisitLogRouter(svcCtx *svctx.ServiceContext) *VisitLogRouter {
	return &VisitLogRouter{
		svcCtx: svcCtx,
	}
}

func (r *VisitLogRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewVisitLogHandler(r.svcCtx)
		// 批量删除访问日志
		g.DELETE("/visit_log/delete_visit_log", h.DeleteVisitLog)
		// 获取访问日志列表
		g.POST("/visit_log/query_visit_log_list", h.QueryVisitLogList)
	}
}
