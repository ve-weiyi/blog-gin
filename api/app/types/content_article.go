package types

type ArticleDetails struct {
	ArticleHome
	Author               *UserInfoVO       `json:"author"`
	LastArticle          *ArticlePreview   `json:"last_article"`
	NextArticle          *ArticlePreview   `json:"next_article"`
	RecommendArticleList []*ArticlePreview `json:"recommend_article_list"`
	NewestArticleList    []*ArticlePreview `json:"newest_article_list"`
}

type GetArticleReq struct {
	ArticleId int64 `form:"article_id"`
}

type LikeArticleReq struct {
	ArticleId int64 `json:"article_id"`
}

// 归档文章查询
type QueryArchivedArticleListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
}

// 列表查询（按你的规范命名）
type QueryArticleListReq struct {
	PageQuery
	Page         *int64    `json:"page,optional"`      // 当前页码
	PageSize     *int64    `json:"page_size,optional"` // 每页数量
	Sorts        *[]string `json:"sorts,optional"`     // 排序
	ArticleTitle *string   `json:"article_title,optional"`
	CategoryName *string   `json:"category_name,optional"`
	TagName      *string   `json:"tag_name,optional"`
}

// 推荐文章查询
type QueryRecommendArticleListReq struct {
}
