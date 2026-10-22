package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type PageRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewPageRouter(svcCtx *svctx.ServiceContext) *PageRouter {
	return &PageRouter{
		svcCtx: svcCtx,
	}
}

func (r *PageRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewPageHandler(r.svcCtx)
		// 创建页面
		g.POST("/page/create_page", h.CreatePage)
		// 删除页面
		g.DELETE("/page/delete_page", h.DeletePage)
		// 获取页面列表
		g.POST("/page/query_page_list", h.QueryPageList)
		// 更新页面
		g.PUT("/page/update_page", h.UpdatePage)
	}
}
