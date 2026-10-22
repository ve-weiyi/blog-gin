package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type UploadLogLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewUploadLogLogic(svcCtx *svctx.ServiceContext) *UploadLogLogic {
	return &UploadLogLogic{
		svcCtx: svcCtx,
	}
}

// 批量删除文件日志
func (l *UploadLogLogic) DeleteUploadLog(reqCtx *request.Context, req *types.DeleteUploadLogReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取文件日志列表
func (l *UploadLogLogic) QueryUploadLogList(reqCtx *request.Context, req *types.QueryUploadLogListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}
