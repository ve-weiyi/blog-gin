package tokenx

import "context"

// Signer 令牌签名与验证接口。
type Signer interface {
	Sign(ctx context.Context, claims TokenClaims) (string, error)

	// Verify 校验令牌并返回解析后的载荷，验签失败返回 ErrTokenExpired 或 ErrTokenInvalid。
	Verify(ctx context.Context, token string) (*TokenClaims, error)
}
