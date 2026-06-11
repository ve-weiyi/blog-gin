package types

// 获取用户详情请求
type GetUserDetailReq struct {
	UserId string `form:"user_id"` // 用户ID (UUID)
}

// 用户详情响应
type GetUserDetailResp struct {
	UserVO
	LastLogin UserLastLogin `json:"last_login"`
}

type OnlineUserListResp struct {
	List []*OnlineUserItem `json:"list"`
}

// 用户列表查询请求
type QueryUserListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
	Mobile   *string   `json:"mobile,optional"`    // 手机号筛选
	Status   *int64    `json:"status,optional"`    // 状态筛选
	Keyword  *string   `json:"keyword,optional"`   // 关键词搜索（昵称/手机号）
}

// 重置用户密码请求
type ResetUserPasswordReq struct {
	UserId      string `json:"user_id"`      // 用户ID (UUID)
	NewPassword string `json:"new_password"` // 新密码
}

// 重置用户密码响应
type ResetUserPasswordResp struct {
	Success bool `json:"success"`
}

// 更新用户角色请求
type UpdateUserRolesReq struct {
	UserId  string  `json:"user_id"`  // 用户ID (UUID)
	RoleIds []int64 `json:"role_ids"` // 角色ID列表
}

// 更新用户角色响应
type UpdateUserRolesResp struct {
	Success bool `json:"success"`
}

// 更新用户状态请求
type UpdateUserStatusReq struct {
	UserId string `json:"user_id"` // 用户ID (UUID)
	Status int64  `json:"status"`  // 状态：0-禁用 1-正常 2-冻结
}

// 更新用户状态响应
type UpdateUserStatusResp struct {
	Success bool `json:"success"`
}
