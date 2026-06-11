package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type MenuRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewMenuRouter(svcCtx *svctx.ServiceContext) *MenuRouter {
	return &MenuRouter{
		svcCtx: svcCtx,
	}
}

func (r *MenuRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewMenuHandler(r.svcCtx)
		// 清空菜单列表
		g.POST("/menu/clean_menu", h.CleanMenu)
		// 创建菜单
		g.POST("/menu/create_menu", h.CreateMenu)
		// 删除菜单
		g.DELETE("/menu/delete_menu", h.DeleteMenu)
		// 获取菜单详情
		g.GET("/menu/get_menu", h.GetMenu)
		// 获取菜单列表
		g.POST("/menu/query_menu_list", h.QueryMenuList)
		// 同步菜单列表
		g.POST("/menu/sync_menu", h.SyncMenu)
		// 更新菜单
		g.PUT("/menu/update_menu", h.UpdateMenu)
	}
}
