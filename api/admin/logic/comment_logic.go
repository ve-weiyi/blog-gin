package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type CommentLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewCommentLogic(svcCtx *svctx.ServiceContext) *CommentLogic {
	return &CommentLogic{
		svcCtx: svcCtx,
	}
}

// 批量删除评论
func (l *CommentLogic) DeleteComment(reqCtx *request.Context, req *types.DeleteCommentReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取评论列表(后台)
func (l *CommentLogic) QueryCommentList(reqCtx *request.Context, req *types.QueryCommentListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 批量更新评论状态
func (l *CommentLogic) UpdateCommentStatus(reqCtx *request.Context, req *types.UpdateCommentStatusReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}
