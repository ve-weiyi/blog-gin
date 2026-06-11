package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type NotifyTemplateHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewNotifyTemplateHandler(svcCtx *svctx.ServiceContext) *NotifyTemplateHandler {
	return &NotifyTemplateHandler{
		svcCtx: svcCtx,
	}
}

// 创建通知模板
// @Param data body types.CreateNotifyTemplateReq true "请求参数"
// @Success 200 {object} types.NotifyTemplateVO "返回信息"
// @Router /admin-api/v1/notify_template/create_notify_template [POST]
func (h *NotifyTemplateHandler) CreateNotifyTemplate(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateNotifyTemplateReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyTemplateLogic(h.svcCtx)
	resp, err := l.CreateNotifyTemplate(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 批量删除通知模板
// @Param data body types.DeleteNotifyTemplateReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/notify_template/delete_notify_template [DELETE]
func (h *NotifyTemplateHandler) DeleteNotifyTemplate(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteNotifyTemplateReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyTemplateLogic(h.svcCtx)
	resp, err := l.DeleteNotifyTemplate(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取通知模板详情
// @Param data body types.GetNotifyTemplateReq true "请求参数"
// @Success 200 {object} types.NotifyTemplateVO "返回信息"
// @Router /admin-api/v1/notify_template/get_notify_template [GET]
func (h *NotifyTemplateHandler) GetNotifyTemplate(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetNotifyTemplateReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyTemplateLogic(h.svcCtx)
	resp, err := l.GetNotifyTemplate(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取通知模板列表
// @Param data body types.QueryNotifyTemplateListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/notify_template/query_notify_template_list [POST]
func (h *NotifyTemplateHandler) QueryNotifyTemplateList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryNotifyTemplateListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyTemplateLogic(h.svcCtx)
	resp, err := l.QueryNotifyTemplateList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新通知模板
// @Param data body types.UpdateNotifyTemplateReq true "请求参数"
// @Success 200 {object} types.NotifyTemplateVO "返回信息"
// @Router /admin-api/v1/notify_template/update_notify_template [PUT]
func (h *NotifyTemplateHandler) UpdateNotifyTemplate(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateNotifyTemplateReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewNotifyTemplateLogic(h.svcCtx)
	resp, err := l.UpdateNotifyTemplate(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
