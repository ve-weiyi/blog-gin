package logic

import (
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type UploadLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewUploadLogic(svcCtx *svctx.ServiceContext) *UploadLogic {
	return &UploadLogic{
		svcCtx: svcCtx,
	}
}

// 获取上传凭证（前端直传）
func (l *UploadLogic) GetUploadToken(reqCtx *request.Context, req *types.GetUploadTokenReq) (resp *types.GetUploadTokenResp, err error) {
	// TODO: implement business logic

	return
}

// 上传文件（服务端上传）
func (l *UploadLogic) UploadFile(reqCtx *request.Context, req *types.UploadFileReq) (resp *types.UploadFileResp, err error) {
	// TODO: implement business logic

	return
}
