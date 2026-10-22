package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type PhotoLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewPhotoLogic(svcCtx *svctx.ServiceContext) *PhotoLogic {
	return &PhotoLogic{
		svcCtx: svcCtx,
	}
}

// 创建照片
func (l *PhotoLogic) CreatePhoto(reqCtx *request.Context, req *types.CreatePhotoReq) (resp *types.PhotoVO, err error) {
	// TODO: implement business logic

	return
}

// 批量删除照片
func (l *PhotoLogic) DeletePhoto(reqCtx *request.Context, req *types.DeletePhotoReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取照片列表
func (l *PhotoLogic) QueryPhotoList(reqCtx *request.Context, req *types.QueryPhotoListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 更新照片
func (l *PhotoLogic) UpdatePhoto(reqCtx *request.Context, req *types.UpdatePhotoReq) (resp *types.PhotoVO, err error) {
	// TODO: implement business logic

	return
}

// 批量更新照片删除状态
func (l *PhotoLogic) UpdatePhotoDelete(reqCtx *request.Context, req *types.UpdatePhotoDeleteReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}
