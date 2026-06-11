package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/app/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type AlbumHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewAlbumHandler(svcCtx *svctx.ServiceContext) *AlbumHandler {
	return &AlbumHandler{
		svcCtx: svcCtx,
	}
}

// 获取相册详情
// @Param data body types.GetAlbumReq true "请求参数"
// @Success 200 {object} types.Album "返回信息"
// @Router /api/v1/album/get_album [GET]
func (h *AlbumHandler) GetAlbum(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetAlbumReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAlbumLogic(h.svcCtx)
	resp, err := l.GetAlbum(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取相册列表
// @Param data body types.QueryAlbumListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/album/query_album_list [POST]
func (h *AlbumHandler) QueryAlbumList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryAlbumListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAlbumLogic(h.svcCtx)
	resp, err := l.QueryAlbumList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取相册下的照片列表
// @Param data body types.QueryAlbumPhotoListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/album/query_album_photo_list [POST]
func (h *AlbumHandler) QueryAlbumPhotoList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryAlbumPhotoListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAlbumLogic(h.svcCtx)
	resp, err := l.QueryAlbumPhotoList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
