package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type FriendRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewFriendRouter(svcCtx *svctx.ServiceContext) *FriendRouter {
	return &FriendRouter{
		svcCtx: svcCtx,
	}
}

func (r *FriendRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/api/v1")

		h := handler.NewFriendHandler(r.svcCtx)
		// 获取友链列表
		g.POST("/friend/query_friend_list", h.QueryFriendList)
	}
}
