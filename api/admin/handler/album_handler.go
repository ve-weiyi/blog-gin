package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/logic"
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type AlbumHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewAlbumHandler(svcCtx *svctx.ServiceContext) *AlbumHandler {
	return &AlbumHandler{
		svcCtx: svcCtx,
	}
}

// 创建相册
// @Param data body types.CreateAlbumReq true "请求参数"
// @Success 200 {object} types.AlbumVO "返回信息"
// @Router /admin-api/v1/album/create_album [POST]
func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateAlbumReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAlbumLogic(h.svcCtx)
	resp, err := l.CreateAlbum(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 批量删除相册
// @Param data body types.DeleteAlbumReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/album/delete_album [DELETE]
func (h *AlbumHandler) DeleteAlbum(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteAlbumReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAlbumLogic(h.svcCtx)
	resp, err := l.DeleteAlbum(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取相册详情
// @Param data body types.GetAlbumReq true "请求参数"
// @Success 200 {object} types.AlbumVO "返回信息"
// @Router /admin-api/v1/album/get_album [GET]
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
// @Router /admin-api/v1/album/query_album_list [POST]
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

// 更新相册
// @Param data body types.UpdateAlbumReq true "请求参数"
// @Success 200 {object} types.AlbumVO "返回信息"
// @Router /admin-api/v1/album/update_album [PUT]
func (h *AlbumHandler) UpdateAlbum(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateAlbumReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAlbumLogic(h.svcCtx)
	resp, err := l.UpdateAlbum(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 批量更新相册删除状态
// @Param data body types.UpdateAlbumDeleteReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/album/update_album_delete [PUT]
func (h *AlbumHandler) UpdateAlbumDelete(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateAlbumDeleteReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewAlbumLogic(h.svcCtx)
	resp, err := l.UpdateAlbumDelete(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
