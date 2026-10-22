package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/logic"
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type VisitLogHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewVisitLogHandler(svcCtx *svctx.ServiceContext) *VisitLogHandler {
	return &VisitLogHandler{
		svcCtx: svcCtx,
	}
}

// 批量删除访问日志
// @Param data body types.DeleteVisitLogReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/visit_log/delete_visit_log [DELETE]
func (h *VisitLogHandler) DeleteVisitLog(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteVisitLogReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewVisitLogLogic(h.svcCtx)
	resp, err := l.DeleteVisitLog(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取访问日志列表
// @Param data body types.QueryVisitLogListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/visit_log/query_visit_log_list [POST]
func (h *VisitLogHandler) QueryVisitLogList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryVisitLogListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewVisitLogLogic(h.svcCtx)
	resp, err := l.QueryVisitLogList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
