package tokenx

import "time"

// Config 令牌管理器配置。
type Config struct {
	Signer          Signer
	Store           Store
	KeyPrefix       string // key 前缀，如 "blog:app:token"，零值为无前缀
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}
