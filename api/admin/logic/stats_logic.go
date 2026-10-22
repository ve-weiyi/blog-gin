package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type StatsLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewStatsLogic(svcCtx *svctx.ServiceContext) *StatsLogic {
	return &StatsLogic{
		svcCtx: svcCtx,
	}
}

// 获取文章分析数据
func (l *StatsLogic) GetArticleAnalytics(reqCtx *request.Context, req *types.EmptyReq) (resp *types.GetArticleAnalyticsResp, err error) {
	// TODO: implement business logic

	return
}

// 获取仪表盘统计数据
func (l *StatsLogic) GetStatsDashboard(reqCtx *request.Context, req *types.EmptyReq) (resp *types.GetStatsDashboardResp, err error) {
	// TODO: implement business logic

	return
}

// 获取用户地理分布
func (l *StatsLogic) GetUserGeoStats(reqCtx *request.Context, req *types.GetUserGeoStatsReq) (resp *types.GetUserGeoStatsResp, err error) {
	// TODO: implement business logic

	return
}

// 获取访客数据趋势
func (l *StatsLogic) GetVisitTrend(reqCtx *request.Context, req *types.GetVisitTrendReq) (resp *types.GetVisitTrendResp, err error) {
	// TODO: implement business logic

	return
}
