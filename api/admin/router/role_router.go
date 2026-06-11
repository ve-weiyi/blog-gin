package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type RoleRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewRoleRouter(svcCtx *svctx.ServiceContext) *RoleRouter {
	return &RoleRouter{
		svcCtx: svcCtx,
	}
}

func (r *RoleRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewRoleHandler(r.svcCtx)
		// 创建角色
		g.POST("/role/create_role", h.CreateRole)
		// 删除角色
		g.DELETE("/role/delete_role", h.DeleteRole)
		// 获取角色详情
		g.GET("/role/get_role", h.GetRole)
		// 查询角色权限配置
		g.GET("/role/get_role_permissions", h.GetRolePermissions)
		// 获取角色列表
		g.POST("/role/query_role_list", h.QueryRoleList)
		// 更新角色信息
		g.PUT("/role/update_role", h.UpdateRole)
		// 更新角色接口权限
		g.PUT("/role/update_role_api_permissions", h.UpdateRoleApiPermissions)
		// 更新角色菜单权限
		g.PUT("/role/update_role_menu_permissions", h.UpdateRoleMenuPermissions)
	}
}
