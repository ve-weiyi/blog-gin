package bizcode

// 业务错误码（非 HTTP 状态码，放在 JSON body 的 code 字段中）
const (
	CodeSuccess = 200

	// 3xx — 流控/安全
	CodeRateLimit  = 300
	CodeTooManyReq = 301
	CodeCaptchaErr = 302
	CodeCodeExpire = 303

	// 4xx — 客户端错误
	CodeBadRequest   = 400
	CodeUnauthorized = 410
	CodeNoPermission = 420
	CodeNotFound     = 430

	// 业务错误
	CodeUserNotFound      = 411
	CodeUserPasswordWrong = 412
	CodeUserDisabled      = 413
	CodeTokenExpired      = 414
	CodeTokenInvalid      = 415
	CodeAccountExist      = 416

	// 5xx — 服务端错误
	CodeServerError = 500
)

var codeMsg = map[int]string{
	CodeSuccess:           "ok",
	CodeRateLimit:         "请求过于频繁，请稍后再试",
	CodeTooManyReq:        "请求过多",
	CodeCaptchaErr:        "验证码错误",
	CodeCodeExpire:        "验证码已过期",
	CodeBadRequest:        "参数错误",
	CodeUnauthorized:      "未登录或登录已过期",
	CodeNoPermission:      "无权限",
	CodeNotFound:          "资源不存在",
	CodeUserNotFound:      "用户不存在",
	CodeUserPasswordWrong: "密码错误",
	CodeUserDisabled:      "用户已被禁用",
	CodeTokenExpired:      "Token已过期",
	CodeTokenInvalid:      "Token无效",
	CodeAccountExist:      "账号已存在",
	CodeServerError:       "服务器内部错误",
}

func Message(code int) string {
	if msg, ok := codeMsg[code]; ok {
		return msg
	}
	return "未知错误"
}
