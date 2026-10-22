package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type LoginLogRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewLoginLogRouter(svcCtx *svctx.ServiceContext) *LoginLogRouter {
	return &LoginLogRouter{
		svcCtx: svcCtx,
	}
}

func (r *LoginLogRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewLoginLogHandler(r.svcCtx)
		// 批量删除登录日志
		g.DELETE("/login_log/delete_login_log", h.DeleteLoginLog)
		// 获取登录日志列表
		g.POST("/login_log/query_login_log_list", h.QueryLoginLogList)
	}
}
