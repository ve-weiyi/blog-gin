package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/logic"
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type OperationLogHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewOperationLogHandler(svcCtx *svctx.ServiceContext) *OperationLogHandler {
	return &OperationLogHandler{
		svcCtx: svcCtx,
	}
}

// 批量删除操作日志
// @Param data body types.DeleteOperationLogReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/operation_log/delete_operation_log [DELETE]
func (h *OperationLogHandler) DeleteOperationLog(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteOperationLogReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewOperationLogLogic(h.svcCtx)
	resp, err := l.DeleteOperationLog(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取操作日志列表
// @Param data body types.QueryOperationLogListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/operation_log/query_operation_log_list [POST]
func (h *OperationLogHandler) QueryOperationLogList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryOperationLogListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewOperationLogLogic(h.svcCtx)
	resp, err := l.QueryOperationLogList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
