package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/app/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type CommentLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewCommentLogic(svcCtx *svctx.ServiceContext) *CommentLogic {
	return &CommentLogic{
		svcCtx: svcCtx,
	}
}

// 获取评论列表
func (l *CommentLogic) QueryCommentList(reqCtx *request.Context, req *types.QueryCommentListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 获取评论回复列表
func (l *CommentLogic) QueryCommentReplyList(reqCtx *request.Context, req *types.QueryCommentReplyListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 获取最新评论列表
func (l *CommentLogic) QueryRecentCommentList(reqCtx *request.Context, req *types.QueryCommentListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 创建评论
func (l *CommentLogic) CreateComment(reqCtx *request.Context, req *types.CreateCommentReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 点赞评论
func (l *CommentLogic) LikeComment(reqCtx *request.Context, req *types.LikeCommentReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 更新评论
func (l *CommentLogic) UpdateComment(reqCtx *request.Context, req *types.UpdateCommentReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}
