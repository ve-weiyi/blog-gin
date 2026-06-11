package types

type CreateNotifyTemplateReq struct {
	Code    string `json:"code"`
	Channel string `json:"channel"`
	Scene   string `json:"scene"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Enabled *int64 `json:"enabled,optional"`
}

type DeleteNotifyTemplateReq struct {
	Ids []int64 `json:"ids"`
}

type GetNotifyTemplateReq struct {
	Id int64 `form:"id"`
}

type NotifyTemplateVO struct {
	Id        *int64 `json:"id,optional"`
	Code      string `json:"code"`
	Channel   string `json:"channel"`
	Scene     string `json:"scene"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Enabled   int64  `json:"enabled"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type QueryNotifyTemplateListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
	Channel  *string   `json:"channel,optional"`
}

type UpdateNotifyTemplateReq struct {
	Id      int64  `json:"id"`
	Code    string `json:"code"`
	Channel string `json:"channel"`
	Scene   string `json:"scene"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Enabled int64  `json:"enabled"`
}
