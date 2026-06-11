package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
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
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewFriendHandler(r.svcCtx)
		// 创建友链
		g.POST("/friend/create_friend", h.CreateFriend)
		// 批量删除友链
		g.DELETE("/friend/delete_friend", h.DeleteFriend)
		// 获取友链列表
		g.POST("/friend/query_friend_list", h.QueryFriendList)
		// 更新友链
		g.PUT("/friend/update_friend", h.UpdateFriend)
	}
}
