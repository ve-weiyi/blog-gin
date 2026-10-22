package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type AlbumRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewAlbumRouter(svcCtx *svctx.ServiceContext) *AlbumRouter {
	return &AlbumRouter{
		svcCtx: svcCtx,
	}
}

func (r *AlbumRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewAlbumHandler(r.svcCtx)
		// 创建相册
		g.POST("/album/create_album", h.CreateAlbum)
		// 批量删除相册
		g.DELETE("/album/delete_album", h.DeleteAlbum)
		// 获取相册详情
		g.GET("/album/get_album", h.GetAlbum)
		// 获取相册列表
		g.POST("/album/query_album_list", h.QueryAlbumList)
		// 更新相册
		g.PUT("/album/update_album", h.UpdateAlbum)
		// 批量更新相册删除状态
		g.PUT("/album/update_album_delete", h.UpdateAlbumDelete)
	}
}
