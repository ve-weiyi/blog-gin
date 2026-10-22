package tokenx

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// key type 常量
const (
	keyTypeAccess  = "access"
	keyTypeRefresh = "refresh"
)

type manager struct {
	signer     Signer
	store      Store
	keyPrefix  string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

// New 创建 Manager 实例。Signer/Store 为必填，nil 时返回 error。
// AccessTokenTTL 零值时兜底 15 分钟，RefreshTokenTTL 零值时兜底 7 天。
func New(cfg Config) (Manager, error) {
	if cfg.Signer == nil {
		return nil, errors.New("tokenx: Signer is required")
	}
	if cfg.Store == nil {
		return nil, errors.New("tokenx: Store is required")
	}
	accessTTL := cfg.AccessTokenTTL
	if accessTTL <= 0 {
		accessTTL = 15 * time.Minute
	}
	refreshTTL := cfg.RefreshTokenTTL
	if refreshTTL <= 0 {
		refreshTTL = 7 * 24 * time.Hour
	}
	return &manager{
		signer:     cfg.Signer,
		store:      cfg.Store,
		keyPrefix:  cfg.KeyPrefix,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}, nil
}

// validateUserID 校验 userID 非空。
func validateUserID(userID string) error {
	if userID == "" {
		return errors.New("tokenx: userID must not be empty")
	}
	return nil
}

// formatKey 构造 Store key：prefix:type:userID，prefix 为空时退化为 type:userID。
func (m *manager) formatKey(keyType, userID string) string {
	if m.keyPrefix == "" {
		return fmt.Sprintf("%s:%s", keyType, userID)
	}
	return fmt.Sprintf("%s:%s:%s", m.keyPrefix, keyType, userID)
}

// verifyToken 通用令牌校验。校验通过返回 nil。
func (m *manager) verifyToken(ctx context.Context, userID, token, keyType string) error {
	if token == "" {
		return ErrTokenEmpty
	}

	claims, err := m.signer.Verify(ctx, token)
	if err != nil {
		return err
	}
	if claims.UserID != userID {
		return ErrTokenInvalid
	}
	if claims.TokenType != "" && claims.TokenType != keyType {
		return ErrTokenInvalid
	}

	// 令牌必须是当前存储中的那一枚，否则已被轮换或吊销
	stored, err := m.store.Get(ctx, m.formatKey(keyType, userID))
	if err != nil {
		return err
	}
	if stored == "" {
		return ErrTokenExpired
	}
	if stored != token {
		return ErrTokenInvalid
	}
	return nil
}

// Generate 生成令牌对。单点语义：先吊销该用户已有令牌，再写入新令牌。
func (m *manager) Generate(ctx context.Context, userID string) (*TokenPair, error) {
	if err := validateUserID(userID); err != nil {
		return nil, err
	}

	now := time.Now()

	accessToken, err := m.signer.Sign(ctx, TokenClaims{
		UserID:    userID,
		TokenType: keyTypeAccess,
		IssuedAt:  now,
		ExpiresAt: now.Add(m.accessTTL),
	})
	if err != nil {
		return nil, err
	}
	refreshToken, err := m.signer.Sign(ctx, TokenClaims{
		UserID:    userID,
		TokenType: keyTypeRefresh,
		IssuedAt:  now,
		ExpiresAt: now.Add(m.refreshTTL),
	})
	if err != nil {
		return nil, err
	}

	// 单点登录：吊销该用户已有令牌
	if err := m.Revoke(ctx, userID); err != nil {
		return nil, err
	}

	accessKey := m.formatKey(keyTypeAccess, userID)
	refreshKey := m.formatKey(keyTypeRefresh, userID)

	if err := m.store.Set(ctx, accessKey, accessToken, m.accessTTL); err != nil {
		return nil, err
	}
	if err := m.store.Set(ctx, refreshKey, refreshToken, m.refreshTTL); err != nil {
		m.store.Delete(ctx, accessKey) // 回滚
		return nil, err
	}

	return &TokenPair{
		UserID:           userID,
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		AccessExpiresAt:  now.Add(m.accessTTL),
		AccessExpiresIn:  int64(m.accessTTL.Seconds()),
		RefreshExpiresAt: now.Add(m.refreshTTL),
		RefreshExpiresIn: int64(m.refreshTTL.Seconds()),
	}, nil
}

// Validate 校验 AccessToken。
func (m *manager) Validate(ctx context.Context, userID, accessToken string) error {
	if err := validateUserID(userID); err != nil {
		return err
	}
	return m.verifyToken(ctx, userID, accessToken, keyTypeAccess)
}

// Refresh 校验 RefreshToken 并轮换出新令牌对。
// Generate 会覆盖 refresh key，旧 RefreshToken 自动失效。
func (m *manager) Refresh(ctx context.Context, userID, refreshToken string) (*TokenPair, error) {
	if err := validateUserID(userID); err != nil {
		return nil, err
	}
	if err := m.verifyToken(ctx, userID, refreshToken, keyTypeRefresh); err != nil {
		return nil, err
	}
	return m.Generate(ctx, userID)
}

// Revoke 吊销该用户的全部令牌。
func (m *manager) Revoke(ctx context.Context, userID string) error {
	if err := validateUserID(userID); err != nil {
		return err
	}
	return m.store.Delete(ctx,
		m.formatKey(keyTypeAccess, userID),
		m.formatKey(keyTypeRefresh, userID),
	)
}
