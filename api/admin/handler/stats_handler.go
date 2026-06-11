package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type StatsHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewStatsHandler(svcCtx *svctx.ServiceContext) *StatsHandler {
	return &StatsHandler{
		svcCtx: svcCtx,
	}
}

// 获取文章分析数据
// @Param data body types.EmptyReq true "请求参数"
// @Success 200 {object} types.GetArticleAnalyticsResp "返回信息"
// @Router /admin-api/v1/stats/get_article_analytics [GET]
func (h *StatsHandler) GetArticleAnalytics(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewStatsLogic(h.svcCtx)
	resp, err := l.GetArticleAnalytics(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取仪表盘统计数据
// @Param data body types.EmptyReq true "请求参数"
// @Success 200 {object} types.GetStatsDashboardResp "返回信息"
// @Router /admin-api/v1/stats/get_stats_dashboard [GET]
func (h *StatsHandler) GetStatsDashboard(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewStatsLogic(h.svcCtx)
	resp, err := l.GetStatsDashboard(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取用户地理分布
// @Param data body types.GetUserGeoStatsReq true "请求参数"
// @Success 200 {object} types.GetUserGeoStatsResp "返回信息"
// @Router /admin-api/v1/stats/get_user_geo_stats [POST]
func (h *StatsHandler) GetUserGeoStats(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetUserGeoStatsReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewStatsLogic(h.svcCtx)
	resp, err := l.GetUserGeoStats(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取访客数据趋势
// @Param data body types.GetVisitTrendReq true "请求参数"
// @Success 200 {object} types.GetVisitTrendResp "返回信息"
// @Router /admin-api/v1/stats/get_visit_trend [POST]
func (h *StatsHandler) GetVisitTrend(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetVisitTrendReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewStatsLogic(h.svcCtx)
	resp, err := l.GetVisitTrend(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
