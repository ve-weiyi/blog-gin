package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
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

// 创建相册
func (l *AlbumLogic) CreateAlbum(reqCtx *request.Context, req *types.CreateAlbumReq) (resp *types.AlbumVO, err error) {
	// TODO: implement business logic

	return
}

// 批量删除相册
func (l *AlbumLogic) DeleteAlbum(reqCtx *request.Context, req *types.DeleteAlbumReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取相册详情
func (l *AlbumLogic) GetAlbum(reqCtx *request.Context, req *types.GetAlbumReq) (resp *types.AlbumVO, err error) {
	// TODO: implement business logic

	return
}

// 获取相册列表
func (l *AlbumLogic) QueryAlbumList(reqCtx *request.Context, req *types.QueryAlbumListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 更新相册
func (l *AlbumLogic) UpdateAlbum(reqCtx *request.Context, req *types.UpdateAlbumReq) (resp *types.AlbumVO, err error) {
	// TODO: implement business logic

	return
}

// 批量更新相册删除状态
func (l *AlbumLogic) UpdateAlbumDelete(reqCtx *request.Context, req *types.UpdateAlbumDeleteReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}
