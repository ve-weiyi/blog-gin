package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type VisitLogLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewVisitLogLogic(svcCtx *svctx.ServiceContext) *VisitLogLogic {
	return &VisitLogLogic{
		svcCtx: svcCtx,
	}
}

// 批量删除访问日志
func (l *VisitLogLogic) DeleteVisitLog(reqCtx *request.Context, req *types.DeleteVisitLogReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取访问日志列表
func (l *VisitLogLogic) QueryVisitLogList(reqCtx *request.Context, req *types.QueryVisitLogListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}
