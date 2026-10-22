package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/handler"
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
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewUploadHandler(r.svcCtx)
		// 获取上传凭证（前端直传）
		g.POST("/upload/get_upload_token", h.GetUploadToken)
		// 查询上传文件列表
		g.POST("/upload/query_file_list", h.QueryFileList)
		// 上传文件（服务端上传）
		g.POST("/upload/upload_file", h.UploadFile)
	}
}
