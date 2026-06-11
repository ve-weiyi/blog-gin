package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type LoginLogHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewLoginLogHandler(svcCtx *svctx.ServiceContext) *LoginLogHandler {
	return &LoginLogHandler{
		svcCtx: svcCtx,
	}
}

// 批量删除登录日志
// @Param data body types.DeleteLoginLogReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/login_log/delete_login_log [DELETE]
func (h *LoginLogHandler) DeleteLoginLog(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteLoginLogReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewLoginLogLogic(h.svcCtx)
	resp, err := l.DeleteLoginLog(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取登录日志列表
// @Param data body types.QueryLoginLogListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/login_log/query_login_log_list [POST]
func (h *LoginLogHandler) QueryLoginLogList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryLoginLogListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewLoginLogLogic(h.svcCtx)
	resp, err := l.QueryLoginLogList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
