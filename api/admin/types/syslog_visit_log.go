package types

type DeleteVisitLogReq struct {
	Ids []int64 `json:"ids"`
}

type QueryVisitLogListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
	UserId   *string   `json:"user_id,optional"`
	DeviceId *string   `json:"device_id,optional"`
	PageName *string   `json:"page_name,optional"`
}
