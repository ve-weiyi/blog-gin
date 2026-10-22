package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/logic"
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type UploadHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewUploadHandler(svcCtx *svctx.ServiceContext) *UploadHandler {
	return &UploadHandler{
		svcCtx: svcCtx,
	}
}

// 获取上传凭证（前端直传）
// @Param data body types.GetUploadTokenReq true "请求参数"
// @Success 200 {object} types.GetUploadTokenResp "返回信息"
// @Router /api/v1/upload/get_upload_token [POST]
func (h *UploadHandler) GetUploadToken(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetUploadTokenReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewUploadLogic(h.svcCtx)
	resp, err := l.GetUploadToken(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 上传文件（服务端上传）
// @Param data body types.UploadFileReq true "请求参数"
// @Success 200 {object} types.UploadFileResp "返回信息"
// @Router /api/v1/upload/upload_file [POST]
func (h *UploadHandler) UploadFile(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UploadFileReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewUploadLogic(h.svcCtx)
	resp, err := l.UploadFile(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
