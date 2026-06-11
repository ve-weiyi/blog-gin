package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type ApiRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewApiRouter(svcCtx *svctx.ServiceContext) *ApiRouter {
	return &ApiRouter{
		svcCtx: svcCtx,
	}
}

func (r *ApiRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewApiHandler(r.svcCtx)
		// 清空接口列表
		g.POST("/api/clean_api", h.CleanApi)
		// 创建接口
		g.POST("/api/create_api", h.CreateApi)
		// 删除接口
		g.DELETE("/api/delete_api", h.DeleteApi)
		// 获取接口详情
		g.GET("/api/get_api", h.GetApi)
		// 获取接口列表
		g.POST("/api/query_api_list", h.QueryApiList)
		// 同步接口列表
		g.POST("/api/sync_api", h.SyncApi)
		// 更新接口
		g.PUT("/api/update_api", h.UpdateApi)
	}
}
