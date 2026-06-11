package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type NotifyTemplateLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewNotifyTemplateLogic(svcCtx *svctx.ServiceContext) *NotifyTemplateLogic {
	return &NotifyTemplateLogic{
		svcCtx: svcCtx,
	}
}

// 创建通知模板
func (l *NotifyTemplateLogic) CreateNotifyTemplate(reqCtx *request.Context, req *types.CreateNotifyTemplateReq) (resp *types.NotifyTemplateVO, err error) {
	// TODO: implement business logic

	return
}

// 批量删除通知模板
func (l *NotifyTemplateLogic) DeleteNotifyTemplate(reqCtx *request.Context, req *types.DeleteNotifyTemplateReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取通知模板详情
func (l *NotifyTemplateLogic) GetNotifyTemplate(reqCtx *request.Context, req *types.GetNotifyTemplateReq) (resp *types.NotifyTemplateVO, err error) {
	// TODO: implement business logic

	return
}

// 获取通知模板列表
func (l *NotifyTemplateLogic) QueryNotifyTemplateList(reqCtx *request.Context, req *types.QueryNotifyTemplateListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 更新通知模板
func (l *NotifyTemplateLogic) UpdateNotifyTemplate(reqCtx *request.Context, req *types.UpdateNotifyTemplateReq) (resp *types.NotifyTemplateVO, err error) {
	// TODO: implement business logic

	return
}
