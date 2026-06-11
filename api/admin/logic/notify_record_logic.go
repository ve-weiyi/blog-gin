package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type NotifyRecordLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewNotifyRecordLogic(svcCtx *svctx.ServiceContext) *NotifyRecordLogic {
	return &NotifyRecordLogic{
		svcCtx: svcCtx,
	}
}

// 批量标记投递记录为已读
func (l *NotifyRecordLogic) BatchMarkRecordsRead(reqCtx *request.Context, req *types.BatchMarkRecordsReadReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 批量删除投递记录
func (l *NotifyRecordLogic) DeleteNotifyRecord(reqCtx *request.Context, req *types.DeleteNotifyRecordReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 全部标记投递记录为已读
func (l *NotifyRecordLogic) MarkAllRecordsRead(reqCtx *request.Context, req *types.MarkAllRecordsReadReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 标记单条投递记录为已读
func (l *NotifyRecordLogic) MarkRecordRead(reqCtx *request.Context, req *types.MarkRecordReadReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取统一投递记录列表
func (l *NotifyRecordLogic) QueryNotifyRecordList(reqCtx *request.Context, req *types.QueryNotifyRecordListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 查询用户 inbox 投递列表
func (l *NotifyRecordLogic) QueryUserInboxRecordList(reqCtx *request.Context, req *types.QueryUserInboxRecordListReq) (resp *types.QueryUserInboxRecordListResp, err error) {
	// TODO: implement business logic

	return
}
