package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type UploadLogHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewUploadLogHandler(svcCtx *svctx.ServiceContext) *UploadLogHandler {
	return &UploadLogHandler{
		svcCtx: svcCtx,
	}
}

// 批量删除文件日志
// @Param data body types.DeleteUploadLogReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/upload_log/delete_upload_log [DELETE]
func (h *UploadLogHandler) DeleteUploadLog(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteUploadLogReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewUploadLogLogic(h.svcCtx)
	resp, err := l.DeleteUploadLog(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取文件日志列表
// @Param data body types.QueryUploadLogListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/upload_log/query_upload_log_list [POST]
func (h *UploadLogHandler) QueryUploadLogList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryUploadLogListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewUploadLogLogic(h.svcCtx)
	resp, err := l.QueryUploadLogList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
