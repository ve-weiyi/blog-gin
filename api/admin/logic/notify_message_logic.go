package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type NotifyMessageLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewNotifyMessageLogic(svcCtx *svctx.ServiceContext) *NotifyMessageLogic {
	return &NotifyMessageLogic{
		svcCtx: svcCtx,
	}
}

// 创建统一通知消息
func (l *NotifyMessageLogic) CreateNotifyMessage(reqCtx *request.Context, req *types.CreateNotifyMessageReq) (resp *types.NotifyMessageVO, err error) {
	// TODO: implement business logic

	return
}

// 批量删除统一通知消息
func (l *NotifyMessageLogic) DeleteNotifyMessage(reqCtx *request.Context, req *types.DeleteNotifyMessageReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取统一通知消息详情
func (l *NotifyMessageLogic) GetNotifyMessage(reqCtx *request.Context, req *types.GetNotifyMessageReq) (resp *types.NotifyMessageVO, err error) {
	// TODO: implement business logic

	return
}

// 发布统一通知消息
func (l *NotifyMessageLogic) PublishNotifyMessage(reqCtx *request.Context, req *types.PublishNotifyMessageReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取统一通知消息列表
func (l *NotifyMessageLogic) QueryNotifyMessageList(reqCtx *request.Context, req *types.QueryNotifyMessageListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 撤回统一通知消息
func (l *NotifyMessageLogic) RevokeNotifyMessage(reqCtx *request.Context, req *types.RevokeNotifyMessageReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 更新统一通知消息
func (l *NotifyMessageLogic) UpdateNotifyMessage(reqCtx *request.Context, req *types.UpdateNotifyMessageReq) (resp *types.NotifyMessageVO, err error) {
	// TODO: implement business logic

	return
}
