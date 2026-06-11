package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type ArticleRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewArticleRouter(svcCtx *svctx.ServiceContext) *ArticleRouter {
	return &ArticleRouter{
		svcCtx: svcCtx,
	}
}

func (r *ArticleRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewArticleHandler(r.svcCtx)
		// 创建文章
		g.POST("/article/create_article", h.CreateArticle)
		// 删除文章
		g.DELETE("/article/delete_article", h.DeleteArticle)
		// 导出文章列表
		g.POST("/article/export_article", h.ExportArticle)
		// 获取文章详情
		g.GET("/article/get_article", h.GetArticle)
		// 获取文章列表
		g.POST("/article/query_article_list", h.QueryArticleList)
		// 更新文章
		g.PUT("/article/update_article", h.UpdateArticle)
		// 更新文章删除状态
		g.PUT("/article/update_article_delete", h.UpdateArticleDelete)
		// 更新文章置顶状态
		g.PUT("/article/update_article_top", h.UpdateArticleTop)
	}
}
