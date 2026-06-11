package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type UserRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewUserRouter(svcCtx *svctx.ServiceContext) *UserRouter {
	return &UserRouter{
		svcCtx: svcCtx,
	}
}

func (r *UserRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewUserHandler(r.svcCtx)
		// 获取在线用户列表
		g.GET("/user/get_online_users", h.GetOnlineUsers)
		// 获取用户详情
		g.GET("/user/get_user", h.GetUserDetail)
		// 获取用户列表
		g.POST("/user/query_user_list", h.QueryUserList)
		// 重置用户密码
		g.POST("/user/reset_user_password", h.ResetUserPassword)
		// 更新用户角色
		g.PUT("/user/update_user_roles", h.UpdateUserRoles)
		// 更新用户状态
		g.PUT("/user/update_user_status", h.UpdateUserStatus)
	}
}
