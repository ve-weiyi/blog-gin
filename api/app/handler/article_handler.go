package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/app/types"
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

// 获取文章详情
// @Param data body types.GetArticleReq true "请求参数"
// @Success 200 {object} types.ArticleDetails "返回信息"
// @Router /api/v1/article/get_article [GET]
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

// 获取归档文章列表
// @Param data body types.QueryArchivedArticleListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/article/query_archived_article_list [POST]
func (h *ArticleHandler) QueryArchivedArticleList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryArchivedArticleListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewArticleLogic(h.svcCtx)
	resp, err := l.QueryArchivedArticleList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取文章列表
// @Param data body types.QueryArticleListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/article/query_article_list [POST]
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

// 获取推荐文章列表
// @Param data body types.QueryRecommendArticleListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/article/query_recommend_article_list [POST]
func (h *ArticleHandler) QueryRecommendArticleList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryRecommendArticleListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewArticleLogic(h.svcCtx)
	resp, err := l.QueryRecommendArticleList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 点赞文章
// @Param data body types.LikeArticleReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /api/v1/article/like_article [POST]
func (h *ArticleHandler) LikeArticle(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.LikeArticleReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewArticleLogic(h.svcCtx)
	resp, err := l.LikeArticle(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
