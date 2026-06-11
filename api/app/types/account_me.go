package types

// 绑定邮箱请求
type BindUserEmailReq struct {
	Email      string `json:"email"`       // 邮箱
	VerifyCode string `json:"verify_code"` // 验证码
}

// 绑定手机号请求
type BindUserPhoneReq struct {
	Mobile     string `json:"mobile"`      // 手机号
	VerifyCode string `json:"verify_code"` // 验证码
}

// 绑定第三方平台请求
type BindUserThirdPartyReq struct {
	Platform string  `json:"platform"`       // 平台
	Code     string  `json:"code"`           // 授权码
	State    *string `json:"state,optional"` // 状态
}

// 停用账号请求（进入冷静期）
type DeactivateAccountReq struct {
	Password string `json:"password"` // 确认密码
	Reason   string `json:"reason"`   // 停用原因
}

type DeactivateAccountResp struct {
	CoolingPeriodDays  int64 `json:"cooling_period_days"`  // 冷静期天数
	CanReactivateUntil int64 `json:"can_reactivate_until"` // 可恢复截止时间（时间戳，毫秒）
}

// 获取用户点赞集合请求
type GetUserLikeReq struct {
}

// 获取用户点赞集合响应
type GetUserLikeResp struct {
	ArticleLikeSet []int64 `json:"article_like_set"` // 文章点赞集合
	CommentLikeSet []int64 `json:"comment_like_set"` // 评论点赞集合
	TalkLikeSet    []int64 `json:"talk_like_set"`    // 说说点赞集合
}

// 获取当前用户信息请求
type GetUserProfileReq struct {
}

// 获取当前用户信息响应
type GetUserProfileResp struct {
	UserInfoExt
	UserId     string                `json:"user_id"`     // 用户ID（UUID）
	Username   string                `json:"username"`    // 用户名
	Nickname   string                `json:"nickname"`    // 昵称
	Avatar     string                `json:"avatar"`      // 头像URL
	Email      string                `json:"email"`       // 邮箱
	Mobile     string                `json:"mobile"`      // 手机号
	Status     int64                 `json:"status"`      // 账号状态
	Plan       string                `json:"plan"`        // 会员计划
	Balance    float64               `json:"balance"`     // 余额
	Coin       int64                 `json:"coin"`        // 账户积分
	CreatedAt  int64                 `json:"created_at"`  // 注册时间
	ThirdParty []*UserThirdPartyInfo `json:"third_party"` // 第三方绑定
}

// 恢复账号请求（冷静期内）
type ReactivateAccountReq struct {
	Email            *string `json:"email,optional"`    // 邮箱
	Mobile           *string `json:"mobile,optional"`   // 手机号
	VerificationCode string  `json:"verification_code"` // 验证码（可选）
}

type ReactivateAccountResp struct {
	Success bool `json:"success"` // 是否恢复成功
}

// 解绑第三方平台请求
type UnbindUserThirdPartyReq struct {
	Platform string `json:"platform"` // 平台
}

// 更新用户信息请求
type UpdateUserProfileReq struct {
	UserInfoExt
	Nickname *string `json:"nickname,optional"` // 昵称
	Avatar   *string `json:"avatar,optional"`   // 头像URL
}

// 更新用户信息响应
type UpdateUserProfileResp struct {
}
