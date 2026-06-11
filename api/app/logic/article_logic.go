package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/app/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type ArticleLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewArticleLogic(svcCtx *svctx.ServiceContext) *ArticleLogic {
	return &ArticleLogic{
		svcCtx: svcCtx,
	}
}

// 获取文章详情
func (l *ArticleLogic) GetArticle(reqCtx *request.Context, req *types.GetArticleReq) (resp *types.ArticleDetails, err error) {
	// TODO: implement business logic

	return
}

// 获取归档文章列表
func (l *ArticleLogic) QueryArchivedArticleList(reqCtx *request.Context, req *types.QueryArchivedArticleListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 获取文章列表
func (l *ArticleLogic) QueryArticleList(reqCtx *request.Context, req *types.QueryArticleListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 获取推荐文章列表
func (l *ArticleLogic) QueryRecommendArticleList(reqCtx *request.Context, req *types.QueryRecommendArticleListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 点赞文章
func (l *ArticleLogic) LikeArticle(reqCtx *request.Context, req *types.LikeArticleReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}
