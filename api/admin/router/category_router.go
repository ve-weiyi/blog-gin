package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type CategoryRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewCategoryRouter(svcCtx *svctx.ServiceContext) *CategoryRouter {
	return &CategoryRouter{
		svcCtx: svcCtx,
	}
}

func (r *CategoryRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewCategoryHandler(r.svcCtx)
		// 创建分类
		g.POST("/category/create_category", h.CreateCategory)
		// 批量删除分类
		g.DELETE("/category/delete_category", h.DeleteCategory)
		// 获取分类列表
		g.POST("/category/query_category_list", h.QueryCategoryList)
		// 更新分类
		g.PUT("/category/update_category", h.UpdateCategory)
	}
}
