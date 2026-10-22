package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/logic"
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type CommentHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewCommentHandler(svcCtx *svctx.ServiceContext) *CommentHandler {
	return &CommentHandler{
		svcCtx: svcCtx,
	}
}

// 批量删除评论
// @Param data body types.DeleteCommentReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/comment/delete_comment [DELETE]
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteCommentReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewCommentLogic(h.svcCtx)
	resp, err := l.DeleteComment(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取评论列表(后台)
// @Param data body types.QueryCommentListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/comment/query_comment_list [POST]
func (h *CommentHandler) QueryCommentList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryCommentListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewCommentLogic(h.svcCtx)
	resp, err := l.QueryCommentList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 批量更新评论状态
// @Param data body types.UpdateCommentStatusReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/comment/update_comment_status [PUT]
func (h *CommentHandler) UpdateCommentStatus(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateCommentStatusReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewCommentLogic(h.svcCtx)
	resp, err := l.UpdateCommentStatus(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
