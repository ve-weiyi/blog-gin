package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/handler"
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
		g := rg.Group("/api/v1")

		h := handler.NewCategoryHandler(r.svcCtx)
		// 获取分类列表
		g.POST("/category/query_category_list", h.QueryCategoryList)
	}
}
