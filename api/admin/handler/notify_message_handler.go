package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type NotifyMessageHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewNotifyMessageHandler(svcCtx *svctx.ServiceContext) *NotifyMessageHandler {
	return &NotifyMessageHandler{
		svcCtx: svcCtx,
	}
}

// 创建统一通知消息
// @Param data body types.CreateNotifyMessageReq true "请求参数"
// @Success 200 {object} types.NotifyMessageVO "返回信息"
// @Router /admin-api/v1/notify_message/create_notify_message [POST]
func (h *NotifyMessageHandler) CreateNotifyMessage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateNotifyMessageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyMessageLogic(h.svcCtx)
	resp, err := l.CreateNotifyMessage(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 批量删除统一通知消息
// @Param data body types.DeleteNotifyMessageReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/notify_message/delete_notify_message [DELETE]
func (h *NotifyMessageHandler) DeleteNotifyMessage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteNotifyMessageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyMessageLogic(h.svcCtx)
	resp, err := l.DeleteNotifyMessage(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取统一通知消息详情
// @Param data body types.GetNotifyMessageReq true "请求参数"
// @Success 200 {object} types.NotifyMessageVO "返回信息"
// @Router /admin-api/v1/notify_message/get_notify_message [GET]
func (h *NotifyMessageHandler) GetNotifyMessage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetNotifyMessageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyMessageLogic(h.svcCtx)
	resp, err := l.GetNotifyMessage(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 发布统一通知消息
// @Param data body types.PublishNotifyMessageReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/notify_message/publish_notify_message [POST]
func (h *NotifyMessageHandler) PublishNotifyMessage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.PublishNotifyMessageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyMessageLogic(h.svcCtx)
	resp, err := l.PublishNotifyMessage(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取统一通知消息列表
// @Param data body types.QueryNotifyMessageListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/notify_message/query_notify_message_list [POST]
func (h *NotifyMessageHandler) QueryNotifyMessageList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryNotifyMessageListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyMessageLogic(h.svcCtx)
	resp, err := l.QueryNotifyMessageList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 撤回统一通知消息
// @Param data body types.RevokeNotifyMessageReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/notify_message/revoke_notify_message [POST]
func (h *NotifyMessageHandler) RevokeNotifyMessage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.RevokeNotifyMessageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyMessageLogic(h.svcCtx)
	resp, err := l.RevokeNotifyMessage(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新统一通知消息
// @Param data body types.UpdateNotifyMessageReq true "请求参数"
// @Success 200 {object} types.NotifyMessageVO "返回信息"
// @Router /admin-api/v1/notify_message/update_notify_message [PUT]
func (h *NotifyMessageHandler) UpdateNotifyMessage(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateNotifyMessageReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyMessageLogic(h.svcCtx)
	resp, err := l.UpdateNotifyMessage(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
