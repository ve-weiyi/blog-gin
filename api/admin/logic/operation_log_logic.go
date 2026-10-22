package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type OperationLogLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewOperationLogLogic(svcCtx *svctx.ServiceContext) *OperationLogLogic {
	return &OperationLogLogic{
		svcCtx: svcCtx,
	}
}

// 批量删除操作日志
func (l *OperationLogLogic) DeleteOperationLog(reqCtx *request.Context, req *types.DeleteOperationLogReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取操作日志列表
func (l *OperationLogLogic) QueryOperationLogList(reqCtx *request.Context, req *types.QueryOperationLogListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}
