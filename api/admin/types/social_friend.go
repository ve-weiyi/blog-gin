package types

type CreateFriendReq struct {
	LinkName    string `json:"link_name"`
	LinkAvatar  string `json:"link_avatar"`
	LinkAddress string `json:"link_address"`
	LinkIntro   string `json:"link_intro"`
}

type DeleteFriendReq struct {
	Ids []int64 `json:"ids"`
}

type FriendVO struct {
	Id          *int64 `json:"id,optional"`
	LinkName    string `json:"link_name"`
	LinkAvatar  string `json:"link_avatar"`
	LinkAddress string `json:"link_address"`
	LinkIntro   string `json:"link_intro"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

type QueryFriendListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
	LinkName *string   `json:"link_name,optional"`
}

type UpdateFriendReq struct {
	Id          int64  `json:"id"`
	LinkName    string `json:"link_name"`
	LinkAvatar  string `json:"link_avatar"`
	LinkAddress string `json:"link_address"`
	LinkIntro   string `json:"link_intro"`
}
