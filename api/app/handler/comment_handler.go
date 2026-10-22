package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/logic"
	"github.com/ve-weiyi/blog-gin/api/app/types"
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

// 获取评论列表
// @Param data body types.QueryCommentListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/comment/query_comment_list [POST]
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

// 获取评论回复列表
// @Param data body types.QueryCommentReplyListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/comment/query_comment_reply_list [POST]
func (h *CommentHandler) QueryCommentReplyList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryCommentReplyListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewCommentLogic(h.svcCtx)
	resp, err := l.QueryCommentReplyList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取最新评论列表
// @Param data body types.QueryCommentListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/comment/query_recent_comment_list [POST]
func (h *CommentHandler) QueryRecentCommentList(c *gin.Context) {
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
	resp, err := l.QueryRecentCommentList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 创建评论
// @Param data body types.CreateCommentReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /api/v1/comment/create_comment [POST]
func (h *CommentHandler) CreateComment(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateCommentReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewCommentLogic(h.svcCtx)
	resp, err := l.CreateComment(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 点赞评论
// @Param data body types.LikeCommentReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /api/v1/comment/like_comment [POST]
func (h *CommentHandler) LikeComment(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.LikeCommentReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewCommentLogic(h.svcCtx)
	resp, err := l.LikeComment(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新评论
// @Param data body types.UpdateCommentReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /api/v1/comment/update_comment [PUT]
func (h *CommentHandler) UpdateComment(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateCommentReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewCommentLogic(h.svcCtx)
	resp, err := l.UpdateComment(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
