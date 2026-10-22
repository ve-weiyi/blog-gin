package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type PhotoRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewPhotoRouter(svcCtx *svctx.ServiceContext) *PhotoRouter {
	return &PhotoRouter{
		svcCtx: svcCtx,
	}
}

func (r *PhotoRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewPhotoHandler(r.svcCtx)
		// 创建照片
		g.POST("/photo/create_photo", h.CreatePhoto)
		// 批量删除照片
		g.DELETE("/photo/delete_photo", h.DeletePhoto)
		// 获取照片列表
		g.POST("/photo/query_photo_list", h.QueryPhotoList)
		// 更新照片
		g.PUT("/photo/update_photo", h.UpdatePhoto)
		// 批量更新照片删除状态
		g.PUT("/photo/update_photo_delete", h.UpdatePhotoDelete)
	}
}
