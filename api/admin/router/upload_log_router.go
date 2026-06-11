package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type UploadLogRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewUploadLogRouter(svcCtx *svctx.ServiceContext) *UploadLogRouter {
	return &UploadLogRouter{
		svcCtx: svcCtx,
	}
}

func (r *UploadLogRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewUploadLogHandler(r.svcCtx)
		// 批量删除文件日志
		g.DELETE("/upload_log/delete_upload_log", h.DeleteUploadLog)
		// 获取文件日志列表
		g.POST("/upload_log/query_upload_log_list", h.QueryUploadLogList)
	}
}
