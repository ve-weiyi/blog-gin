package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
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
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewCommentHandler(r.svcCtx)
		// 批量删除评论
		g.DELETE("/comment/delete_comment", h.DeleteComment)
		// 获取评论列表(后台)
		g.POST("/comment/query_comment_list", h.QueryCommentList)
		// 批量更新评论状态
		g.PUT("/comment/update_comment_status", h.UpdateCommentStatus)
	}
}
