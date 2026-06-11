package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type CommentRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewCommentRouter(svcCtx *svctx.ServiceContext) *CommentRouter {
	return &CommentRouter{
		svcCtx: svcCtx,
	}
}

func (r *CommentRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/api/v1")

		h := handler.NewCommentHandler(r.svcCtx)
		// 获取评论列表
		g.POST("/comment/query_comment_list", h.QueryCommentList)
		// 获取评论回复列表
		g.POST("/comment/query_comment_reply_list", h.QueryCommentReplyList)
		// 获取最新评论列表
		g.POST("/comment/query_recent_comment_list", h.QueryRecentCommentList)
	}
	{
		g := rg.Group("/api/v1")
		g.Use(r.svcCtx.UserAuth)

		h := handler.NewCommentHandler(r.svcCtx)
		// 创建评论
		g.POST("/comment/create_comment", h.CreateComment)
		// 点赞评论
		g.POST("/comment/like_comment", h.LikeComment)
		// 更新评论
		g.PUT("/comment/update_comment", h.UpdateComment)
	}
}
