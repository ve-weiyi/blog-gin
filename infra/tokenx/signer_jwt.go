package tokenx

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtSigner struct {
	key    []byte
	method jwt.SigningMethod
	issuer string
}

// NewJWTSigner 创建 JWT HS256 签名器，生成自含令牌。
func NewJWTSigner(key []byte, issuer string) Signer {
	return &jwtSigner{
		key:    key,
		method: jwt.SigningMethodHS256,
		issuer: issuer,
	}
}

func (s *jwtSigner) Sign(_ context.Context, claims TokenClaims) (string, error) {
	// jti 取随机值：iat/exp 为秒级精度，同一秒内签发的两份相同 claims 若不带 jti 会生成
	// 完全相同的令牌串，导致轮换与吊销失效（旧令牌与新令牌逐字节相等）。
	jwtClaims := jwt.MapClaims{
		"sub": claims.UserID,
		"jti": uuid.NewString(),
		"typ": claims.TokenType,
		"iat": claims.IssuedAt.Unix(),
		"exp": claims.ExpiresAt.Unix(),
		"iss": s.issuer,
	}
	token := jwt.NewWithClaims(s.method, jwtClaims)
	return token.SignedString(s.key)
}

func (s *jwtSigner) Verify(_ context.Context, rawToken string) (*TokenClaims, error) {
	token, err := jwt.Parse(rawToken, func(t *jwt.Token) (interface{}, error) {
		// 防算法混淆攻击
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return s.key, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrTokenInvalid
	}

	uid, _ := claims["sub"].(string)
	tokenType, _ := claims["typ"].(string)
	iatUnix, _ := claims["iat"].(float64)
	expUnix, _ := claims["exp"].(float64)

	return &TokenClaims{
		UserID:    uid,
		TokenType: tokenType,
		IssuedAt:  time.Unix(int64(iatUnix), 0),
		ExpiresAt: time.Unix(int64(expUnix), 0),
	}, nil
}
