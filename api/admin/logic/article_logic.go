package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
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

// 创建文章
func (l *ArticleLogic) CreateArticle(reqCtx *request.Context, req *types.CreateArticleReq) (resp *types.ArticleVO, err error) {
	// TODO: implement business logic

	return
}

// 删除文章
func (l *ArticleLogic) DeleteArticle(reqCtx *request.Context, req *types.DeleteArticleReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 导出文章列表
func (l *ArticleLogic) ExportArticle(reqCtx *request.Context, req *types.ExportArticleReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 获取文章详情
func (l *ArticleLogic) GetArticle(reqCtx *request.Context, req *types.GetArticleReq) (resp *types.ArticleVO, err error) {
	// TODO: implement business logic

	return
}

// 获取文章列表
func (l *ArticleLogic) QueryArticleList(reqCtx *request.Context, req *types.QueryArticleListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 更新文章
func (l *ArticleLogic) UpdateArticle(reqCtx *request.Context, req *types.UpdateArticleReq) (resp *types.ArticleVO, err error) {
	// TODO: implement business logic

	return
}

// 更新文章删除状态
func (l *ArticleLogic) UpdateArticleDelete(reqCtx *request.Context, req *types.UpdateArticleDeleteReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 更新文章置顶状态
func (l *ArticleLogic) UpdateArticleTop(reqCtx *request.Context, req *types.UpdateArticleTopReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}
