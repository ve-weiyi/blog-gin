package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/handler"
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
		g := rg.Group("/api/v1")

		h := handler.NewPageHandler(r.svcCtx)
		// 获取页面列表
		g.POST("/page/query_page_list", h.QueryPageList)
	}
}
