package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type TagHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewTagHandler(svcCtx *svctx.ServiceContext) *TagHandler {
	return &TagHandler{
		svcCtx: svcCtx,
	}
}

// 创建标签
// @Param data body types.CreateTagReq true "请求参数"
// @Success 200 {object} types.TagVO "返回信息"
// @Router /admin-api/v1/tag/create_tag [POST]
func (h *TagHandler) CreateTag(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateTagReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewTagLogic(h.svcCtx)
	resp, err := l.CreateTag(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 删除标签
// @Param data body types.DeleteTagReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/tag/delete_tag [DELETE]
func (h *TagHandler) DeleteTag(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteTagReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewTagLogic(h.svcCtx)
	resp, err := l.DeleteTag(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取标签列表
// @Param data body types.QueryTagListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/tag/query_tag_list [POST]
func (h *TagHandler) QueryTagList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryTagListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewTagLogic(h.svcCtx)
	resp, err := l.QueryTagList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新标签
// @Param data body types.UpdateTagReq true "请求参数"
// @Success 200 {object} types.TagVO "返回信息"
// @Router /admin-api/v1/tag/update_tag [PUT]
func (h *TagHandler) UpdateTag(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateTagReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewTagLogic(h.svcCtx)
	resp, err := l.UpdateTag(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
