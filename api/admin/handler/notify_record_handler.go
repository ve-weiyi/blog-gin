package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type NotifyRecordHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewNotifyRecordHandler(svcCtx *svctx.ServiceContext) *NotifyRecordHandler {
	return &NotifyRecordHandler{
		svcCtx: svcCtx,
	}
}

// 批量标记投递记录为已读
// @Param data body types.BatchMarkRecordsReadReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/notify_record/batch_mark_records_read [POST]
func (h *NotifyRecordHandler) BatchMarkRecordsRead(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.BatchMarkRecordsReadReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyRecordLogic(h.svcCtx)
	resp, err := l.BatchMarkRecordsRead(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 批量删除投递记录
// @Param data body types.DeleteNotifyRecordReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/notify_record/delete_notify_record [DELETE]
func (h *NotifyRecordHandler) DeleteNotifyRecord(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteNotifyRecordReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyRecordLogic(h.svcCtx)
	resp, err := l.DeleteNotifyRecord(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 全部标记投递记录为已读
// @Param data body types.MarkAllRecordsReadReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/notify_record/mark_all_records_read [POST]
func (h *NotifyRecordHandler) MarkAllRecordsRead(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.MarkAllRecordsReadReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyRecordLogic(h.svcCtx)
	resp, err := l.MarkAllRecordsRead(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 标记单条投递记录为已读
// @Param data body types.MarkRecordReadReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/notify_record/mark_record_read [POST]
func (h *NotifyRecordHandler) MarkRecordRead(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.MarkRecordReadReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyRecordLogic(h.svcCtx)
	resp, err := l.MarkRecordRead(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取统一投递记录列表
// @Param data body types.QueryNotifyRecordListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/notify_record/query_notify_record_list [POST]
func (h *NotifyRecordHandler) QueryNotifyRecordList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryNotifyRecordListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyRecordLogic(h.svcCtx)
	resp, err := l.QueryNotifyRecordList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 查询用户 inbox 投递列表
// @Param data body types.QueryUserInboxRecordListReq true "请求参数"
// @Success 200 {object} types.QueryUserInboxRecordListResp "返回信息"
// @Router /admin-api/v1/notify_record/query_user_inbox_record_list [POST]
func (h *NotifyRecordHandler) QueryUserInboxRecordList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryUserInboxRecordListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyRecordLogic(h.svcCtx)
	resp, err := l.QueryUserInboxRecordList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
