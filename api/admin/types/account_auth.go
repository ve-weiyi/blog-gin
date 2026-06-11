package types

// 邮箱验证码登录（仅登录，未注册报错）
type EmailLoginReq struct {
	Email string `json:"email"` // 邮箱
	Code  string `json:"code"`  // 验证码
}

// 邮箱注册（必须设密码）
type EmailRegisterReq struct {
	Email    string  `json:"email"`             // 邮箱
	Password string  `json:"password"`          // 密码
	Code     string  `json:"code"`              // 验证码
	Username *string `json:"username,optional"` // 用户名
	Nickname *string `json:"nickname,optional"` // 昵称
}

type EmailRegisterResp struct {
}

type GetCaptchaReq struct {
	Width  *int64 `form:"width,optional"`  // 宽度
	Height *int64 `form:"height,optional"` // 高度
}

type GetCaptchaResp struct {
	CaptchaKey    string `json:"captcha_key"`    // 验证码key
	CaptchaBase64 string `json:"captcha_base64"` // 验证码base64
	CaptchaCode   string `json:"captcha_code"`   // 验证码
}

// 第三方登录授权URL
type GetOauthAuthorizeUrlReq struct {
	Platform string  `uri:"platform"`        // 平台
	State    *string `form:"state,optional"` // 状态
}

type GetOauthAuthorizeUrlResp struct {
	AuthorizeUrl string `json:"authorize_url"` // 授权地址
}

// 登录响应
type LoginResp struct {
	UserId   string `json:"user_id"`   // 用户id
	UserType string `json:"user_type"` // 用户类型：user-普通用户 admin-管理员
	Scope    string `json:"scope"`     // 作用域
	Token    *Token `json:"token"`
}

type LogoutReq struct {
}

type LogoutResp struct {
}

// 手机验证码登录（自动注册）
type MobileLoginReq struct {
	Mobile string `json:"mobile"` // 手机号
	Code   string `json:"code"`   // 验证码
}

// 第三方登录（前端携带code）
type OauthLoginReq struct {
	Platform string  `uri:"platform"`        // 平台
	Code     string  `json:"code"`           // 授权码
	State    *string `json:"state,optional"` // 状态
}

// 密码登录（账号/手机号/邮箱 + 密码）
type PasswordLoginReq struct {
	Account     string  `json:"account"`               // 账号/手机号/邮箱
	Password    string  `json:"password"`              // 密码
	CaptchaKey  *string `json:"captcha_key,optional"`  // 图形验证码key
	CaptchaCode *string `json:"captcha_code,optional"` // 图形验证码
}

type RefreshTokenReq struct {
	UserId       string `json:"user_id"`       // 用户id
	GrantType    string `json:"grant_type"`    // 授权类型
	RefreshToken string `json:"refresh_token"` // 刷新令牌
}

// 重置密码请求（免登录，通过验证码）
type ResetPasswordReq struct {
	Email           string `json:"email"`            // 邮箱
	Password        string `json:"password"`         // 新密码
	ConfirmPassword string `json:"confirm_password"` // 确认密码
	Code            string `json:"code"`             // 验证码
}

type ResetPasswordResp struct {
}

// 发送邮箱验证码
type SendEmailCodeReq struct {
	Email string `json:"email"` // 邮箱
	Type  string `json:"type"`  // login / register / reset_password / bind_email
}

type SendEmailCodeResp struct {
}

// 发送手机验证码
type SendMobileCodeReq struct {
	Mobile string `json:"mobile"` // 手机号
	Type   string `json:"type"`   // login / reset_password / bind_phone
}

type SendMobileCodeResp struct {
}
