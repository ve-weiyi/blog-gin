package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
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
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewTagHandler(r.svcCtx)
		// 创建标签
		g.POST("/tag/create_tag", h.CreateTag)
		// 删除标签
		g.DELETE("/tag/delete_tag", h.DeleteTag)
		// 获取标签列表
		g.POST("/tag/query_tag_list", h.QueryTagList)
		// 更新标签
		g.PUT("/tag/update_tag", h.UpdateTag)
	}
}
