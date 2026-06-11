package types

type DeleteUploadLogReq struct {
	Ids []int64 `json:"ids"`
}

type QueryUploadLogListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
	FileBase *string   `json:"file_base,optional"`
	FileName *string   `json:"file_name,optional"`
	FileType *string   `json:"file_type,optional"`
}
