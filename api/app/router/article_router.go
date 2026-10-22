package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
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
		g := rg.Group("/api/v1")

		h := handler.NewArticleHandler(r.svcCtx)
		// 获取文章详情
		g.GET("/article/get_article", h.GetArticle)
		// 获取归档文章列表
		g.POST("/article/query_archived_article_list", h.QueryArchivedArticleList)
		// 获取文章列表
		g.POST("/article/query_article_list", h.QueryArticleList)
		// 获取推荐文章列表
		g.POST("/article/query_recommend_article_list", h.QueryRecommendArticleList)
	}
	{
		g := rg.Group("/api/v1")
		g.Use(r.svcCtx.UserAuth)

		h := handler.NewArticleHandler(r.svcCtx)
		// 点赞文章
		g.POST("/article/like_article", h.LikeArticle)
	}
}
