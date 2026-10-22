package logic

import (
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type AlbumLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewAlbumLogic(svcCtx *svctx.ServiceContext) *AlbumLogic {
	return &AlbumLogic{
		svcCtx: svcCtx,
	}
}

// 获取相册详情
func (l *AlbumLogic) GetAlbum(reqCtx *request.Context, req *types.GetAlbumReq) (resp *types.Album, err error) {
	// TODO: implement business logic

	return
}

// 获取相册列表
func (l *AlbumLogic) QueryAlbumList(reqCtx *request.Context, req *types.QueryAlbumListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 获取相册下的照片列表
func (l *AlbumLogic) QueryAlbumPhotoList(reqCtx *request.Context, req *types.QueryAlbumPhotoListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}
