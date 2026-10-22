package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type UploadRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewUploadRouter(svcCtx *svctx.ServiceContext) *UploadRouter {
	return &UploadRouter{
		svcCtx: svcCtx,
	}
}

func (r *UploadRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/api/v1")
		g.Use(r.svcCtx.UserAuth)

		h := handler.NewUploadHandler(r.svcCtx)
		// 获取上传凭证（前端直传）
		g.POST("/upload/get_upload_token", h.GetUploadToken)
		// 上传文件（服务端上传）
		g.POST("/upload/upload_file", h.UploadFile)
	}
}
