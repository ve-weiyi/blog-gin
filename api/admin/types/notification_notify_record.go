package types

type BatchMarkRecordsReadReq struct {
	Ids []int64 `json:"ids"`
}

type DeleteNotifyRecordReq struct {
	Ids []int64 `json:"ids"`
}

type MarkAllRecordsReadReq struct {
	UserId string `json:"user_id"`
}

type MarkRecordReadReq struct {
	Id int64 `json:"id"`
}

type QueryNotifyRecordListReq struct {
	PageQuery
	Page      *int64    `json:"page,optional"`      // 当前页码
	PageSize  *int64    `json:"page_size,optional"` // 每页数量
	Sorts     *[]string `json:"sorts,optional"`     // 排序
	Channel   *string   `json:"channel,optional"`
	Status    *string   `json:"status,optional"`
	Recipient *string   `json:"recipient,optional"`
}

type QueryUserInboxRecordListReq struct {
	PageQuery
	Page       *int64    `json:"page,optional"`      // 当前页码
	PageSize   *int64    `json:"page_size,optional"` // 每页数量
	Sorts      *[]string `json:"sorts,optional"`     // 排序
	UserId     string    `json:"user_id"`
	OnlyUnread *int64    `json:"only_unread,optional"`
	Title      *string   `json:"title,optional"`
}

type QueryUserInboxRecordListResp struct {
	Page        int64       `json:"page"`
	PageSize    int64       `json:"page_size"`
	Total       int64       `json:"total"`
	UnreadTotal int64       `json:"unread_total"`
	List        interface{} `json:"list"`
}
