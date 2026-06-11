package types

type CreateTagReq struct {
	TagName string `json:"tag_name"`
}

type DeleteTagReq struct {
	Ids []int64 `json:"ids"`
}

type QueryTagListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
	TagName  *string   `json:"tag_name,optional"`
}

type TagVO struct {
	Id           *int64 `json:"id,optional"`
	TagName      string `json:"tag_name"`
	ArticleCount int64  `json:"article_count"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type UpdateTagReq struct {
	Id      int64  `json:"id"`
	TagName string `json:"tag_name"`
}
