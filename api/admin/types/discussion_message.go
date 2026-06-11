package types

type DeleteMessageReq struct {
	Ids []int64 `json:"ids"`
}

type QueryMessageListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
	UserId   *string   `json:"user_id,optional"`
	Status   *int64    `json:"status,optional"`
}

type UpdateMessageStatusReq struct {
	Ids    []int64 `json:"ids"`
	Status int64   `json:"status"`
}
