package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/handler"
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
		g := rg.Group("/api/v1")

		h := handler.NewAlbumHandler(r.svcCtx)
		// 获取相册详情
		g.GET("/album/get_album", h.GetAlbum)
		// 获取相册列表
		g.POST("/album/query_album_list", h.QueryAlbumList)
		// 获取相册下的照片列表
		g.POST("/album/query_album_photo_list", h.QueryAlbumPhotoList)
	}
}
