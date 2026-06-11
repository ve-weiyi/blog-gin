package types

type QueryFriendListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
}
