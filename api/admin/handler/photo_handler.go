package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/admin/logic"
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type PhotoHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewPhotoHandler(svcCtx *svctx.ServiceContext) *PhotoHandler {
	return &PhotoHandler{
		svcCtx: svcCtx,
	}
}

// 创建照片
// @Param data body types.CreatePhotoReq true "请求参数"
// @Success 200 {object} types.PhotoVO "返回信息"
// @Router /admin-api/v1/photo/create_photo [POST]
func (h *PhotoHandler) CreatePhoto(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreatePhotoReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewPhotoLogic(h.svcCtx)
	resp, err := l.CreatePhoto(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 批量删除照片
// @Param data body types.DeletePhotoReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/photo/delete_photo [DELETE]
func (h *PhotoHandler) DeletePhoto(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeletePhotoReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewPhotoLogic(h.svcCtx)
	resp, err := l.DeletePhoto(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取照片列表
// @Param data body types.QueryPhotoListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/photo/query_photo_list [POST]
func (h *PhotoHandler) QueryPhotoList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryPhotoListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewPhotoLogic(h.svcCtx)
	resp, err := l.QueryPhotoList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新照片
// @Param data body types.UpdatePhotoReq true "请求参数"
// @Success 200 {object} types.PhotoVO "返回信息"
// @Router /admin-api/v1/photo/update_photo [PUT]
func (h *PhotoHandler) UpdatePhoto(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdatePhotoReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewPhotoLogic(h.svcCtx)
	resp, err := l.UpdatePhoto(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 批量更新照片删除状态
// @Param data body types.UpdatePhotoDeleteReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/photo/update_photo_delete [PUT]
func (h *PhotoHandler) UpdatePhotoDelete(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdatePhotoDeleteReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewPhotoLogic(h.svcCtx)
	resp, err := l.UpdatePhotoDelete(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
