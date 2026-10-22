package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type TagRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewTagRouter(svcCtx *svctx.ServiceContext) *TagRouter {
	return &TagRouter{
		svcCtx: svcCtx,
	}
}

func (r *TagRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/api/v1")

		h := handler.NewTagHandler(r.svcCtx)
		// 获取标签列表
		g.POST("/tag/query_tag_list", h.QueryTagList)
	}
}
