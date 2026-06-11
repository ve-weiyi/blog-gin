package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type ArticleHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewArticleHandler(svcCtx *svctx.ServiceContext) *ArticleHandler {
	return &ArticleHandler{
		svcCtx: svcCtx,
	}
}

// 创建文章
// @Param data body types.CreateArticleReq true "请求参数"
// @Success 200 {object} types.ArticleVO "返回信息"
// @Router /admin-api/v1/article/create_article [POST]
func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateArticleReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewArticleLogic(h.svcCtx)
	resp, err := l.CreateArticle(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 删除文章
// @Param data body types.DeleteArticleReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/article/delete_article [DELETE]
func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteArticleReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewArticleLogic(h.svcCtx)
	resp, err := l.DeleteArticle(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 导出文章列表
// @Param data body types.ExportArticleReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/article/export_article [POST]
func (h *ArticleHandler) ExportArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.ExportArticleReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewArticleLogic(h.svcCtx)
	resp, err := l.ExportArticle(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取文章详情
// @Param data body types.GetArticleReq true "请求参数"
// @Success 200 {object} types.ArticleVO "返回信息"
// @Router /admin-api/v1/article/get_article [GET]
func (h *ArticleHandler) GetArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetArticleReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewArticleLogic(h.svcCtx)
	resp, err := l.GetArticle(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取文章列表
// @Param data body types.QueryArticleListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/article/query_article_list [POST]
func (h *ArticleHandler) QueryArticleList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryArticleListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewArticleLogic(h.svcCtx)
	resp, err := l.QueryArticleList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新文章
// @Param data body types.UpdateArticleReq true "请求参数"
// @Success 200 {object} types.ArticleVO "返回信息"
// @Router /admin-api/v1/article/update_article [PUT]
func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateArticleReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewArticleLogic(h.svcCtx)
	resp, err := l.UpdateArticle(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新文章删除状态
// @Param data body types.UpdateArticleDeleteReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/article/update_article_delete [PUT]
func (h *ArticleHandler) UpdateArticleDelete(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateArticleDeleteReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewArticleLogic(h.svcCtx)
	resp, err := l.UpdateArticleDelete(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新文章置顶状态
// @Param data body types.UpdateArticleTopReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/article/update_article_top [PUT]
func (h *ArticleHandler) UpdateArticleTop(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateArticleTopReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewArticleLogic(h.svcCtx)
	resp, err := l.UpdateArticleTop(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
