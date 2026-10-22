package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type StatsRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewStatsRouter(svcCtx *svctx.ServiceContext) *StatsRouter {
	return &StatsRouter{
		svcCtx: svcCtx,
	}
}

func (r *StatsRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewStatsHandler(r.svcCtx)
		// 获取文章分析数据
		g.GET("/stats/get_article_analytics", h.GetArticleAnalytics)
		// 获取仪表盘统计数据
		g.GET("/stats/get_stats_dashboard", h.GetStatsDashboard)
		// 获取用户地理分布
		g.POST("/stats/get_user_geo_stats", h.GetUserGeoStats)
		// 获取访客数据趋势
		g.POST("/stats/get_visit_trend", h.GetVisitTrend)
	}
}
