package types

type CreateNotifyMessageReq struct {
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Category   string  `json:"category"`
	Level      string  `json:"level"`
	TargetType string  `json:"target_type"`
	TargetIds  *string `json:"target_ids,optional"`
}

type DeleteNotifyMessageReq struct {
	Ids []int64 `json:"ids"`
}

type GetNotifyMessageReq struct {
	Id int64 `form:"id"`
}

type NotifyMessageVO struct {
	Id          *int64 `json:"id,optional"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
	Level       string `json:"level"`
	TargetType  string `json:"target_type"`
	TargetIds   string `json:"target_ids"`
	Status      string `json:"status"`
	PublishedAt int64  `json:"published_at"`
	PublishedBy string `json:"published_by"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type PublishNotifyMessageReq struct {
	Id int64 `json:"id"`
}

type QueryNotifyMessageListReq struct {
	PageQuery
	Page       *int64    `json:"page,optional"`      // 当前页码
	PageSize   *int64    `json:"page_size,optional"` // 每页数量
	Sorts      *[]string `json:"sorts,optional"`     // 排序
	Category   *string   `json:"category,optional"`
	Level      *string   `json:"level,optional"`
	Status     *string   `json:"status,optional"`
	TargetType *string   `json:"target_type,optional"`
}

type RevokeNotifyMessageReq struct {
	Id int64 `json:"id"`
}

type UpdateNotifyMessageReq struct {
	Id         int64   `json:"id"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Category   string  `json:"category"`
	Level      string  `json:"level"`
	TargetType string  `json:"target_type"`
	TargetIds  *string `json:"target_ids,optional"`
}
