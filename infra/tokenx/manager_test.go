package tokenx_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ve-weiyi/blog-gin/infra/tokenx"
)

// memStore Store 的内存实现，仅用于测试。
// 不模拟 TTL：过期行为由 Redis 承担，此处只验证键值与轮换语义。
type memStore struct {
	data map[string]string
}

func newMemStore() *memStore {
	return &memStore{data: make(map[string]string)}
}

func (m *memStore) Set(_ context.Context, key, value string, _ time.Duration) error {
	m.data[key] = value
	return nil
}

func (m *memStore) Get(_ context.Context, key string) (string, error) {
	return m.data[key], nil
}

func (m *memStore) Delete(_ context.Context, keys ...string) error {
	for _, k := range keys {
		delete(m.data, k)
	}
	return nil
}

func newManager(t *testing.T, cfg tokenx.Config) tokenx.Manager {
	t.Helper()
	if cfg.Signer == nil {
		cfg.Signer = tokenx.NewJWTSigner([]byte("test-secret"), "test-issuer")
	}
	if cfg.Store == nil {
		cfg.Store = newMemStore()
	}
	m, err := tokenx.New(cfg)
	if err != nil {
		t.Fatalf("new manager: %v", err)
	}
	return m
}

func TestNew(t *testing.T) {
	t.Run("nil Signer", func(t *testing.T) {
		if _, err := tokenx.New(tokenx.Config{Store: newMemStore()}); err == nil {
			t.Fatal("expected error for nil Signer")
		}
	})

	t.Run("nil Store", func(t *testing.T) {
		cfg := tokenx.Config{Signer: tokenx.NewJWTSigner([]byte("k"), "iss")}
		if _, err := tokenx.New(cfg); err == nil {
			t.Fatal("expected error for nil Store")
		}
	})

	// TTL 零值时兜底：Access 15 分钟，Refresh 7 天。
	t.Run("default TTLs", func(t *testing.T) {
		pair, err := newManager(t, tokenx.Config{}).Generate(context.Background(), "u1")
		if err != nil {
			t.Fatalf("Generate: %v", err)
		}
		if want := int64((15 * time.Minute).Seconds()); pair.AccessExpiresIn != want {
			t.Errorf("AccessExpiresIn = %d, want %d", pair.AccessExpiresIn, want)
		}
		if want := int64((7 * 24 * time.Hour).Seconds()); pair.RefreshExpiresIn != want {
			t.Errorf("RefreshExpiresIn = %d, want %d", pair.RefreshExpiresIn, want)
		}
	})
}

func TestGenerate(t *testing.T) {
	t.Run("returns token pair", func(t *testing.T) {
		pair, err := newManager(t, tokenx.Config{}).Generate(context.Background(), "u1")
		if err != nil {
			t.Fatalf("Generate: %v", err)
		}
		if pair.UserID != "u1" {
			t.Errorf("UserID = %q, want u1", pair.UserID)
		}
		if pair.AccessToken == "" || pair.RefreshToken == "" || pair.AccessToken == pair.RefreshToken {
			t.Errorf("令牌对不合法: %+v", pair)
		}
	})

	t.Run("empty user ID returns error", func(t *testing.T) {
		if _, err := newManager(t, tokenx.Config{}).Generate(context.Background(), ""); err == nil {
			t.Fatal("expected error for empty userID")
		}
	})

	// 单点语义：再次登录后旧令牌立即失效。
	t.Run("revokes previous token", func(t *testing.T) {
		m := newManager(t, tokenx.Config{})
		first, err := m.Generate(context.Background(), "u1")
		if err != nil {
			t.Fatalf("Generate(first): %v", err)
		}
		second, err := m.Generate(context.Background(), "u1")
		if err != nil {
			t.Fatalf("Generate(second): %v", err)
		}

		if err := m.Validate(context.Background(), "u1", first.AccessToken); !errors.Is(err, tokenx.ErrTokenInvalid) {
			t.Errorf("旧 access token = %v, want ErrTokenInvalid", err)
		}
		if err := m.Validate(context.Background(), "u1", second.AccessToken); err != nil {
			t.Errorf("新 access token 校验失败: %v", err)
		}
	})
}

func TestValidate(t *testing.T) {
	ctx := context.Background()
	m := newManager(t, tokenx.Config{})

	pair, err := m.Generate(ctx, "u1")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	// 为 u2 也签一对，确保 u2 在存储中存在，从而走到 claims 比对而非存储缺失。
	if _, err := m.Generate(ctx, "u2"); err != nil {
		t.Fatalf("Generate(u2): %v", err)
	}

	cases := []struct {
		name  string
		user  string
		token string
		want  error
	}{
		{"有效令牌", "u1", pair.AccessToken, nil},
		{"空令牌", "u1", "", tokenx.ErrTokenEmpty},
		{"格式错误", "u1", "not.a.jwt", tokenx.ErrTokenInvalid},
		{"他人令牌", "u2", pair.AccessToken, tokenx.ErrTokenInvalid},
		{"refresh token 不能当 access", "u1", pair.RefreshToken, tokenx.ErrTokenInvalid},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := m.Validate(ctx, tc.user, tc.token)
			if tc.want == nil && err != nil {
				t.Errorf("Validate() = %v, want nil", err)
			}
			if tc.want != nil && !errors.Is(err, tc.want) {
				t.Errorf("Validate() = %v, want %v", err, tc.want)
			}
		})
	}

	t.Run("revoke 后失效", func(t *testing.T) {
		revoked := newManager(t, tokenx.Config{})
		p, err := revoked.Generate(ctx, "u1")
		if err != nil {
			t.Fatalf("Generate: %v", err)
		}
		if err := revoked.Revoke(ctx, "u1"); err != nil {
			t.Fatalf("Revoke: %v", err)
		}
		if err := revoked.Validate(ctx, "u1", p.AccessToken); !errors.Is(err, tokenx.ErrTokenExpired) {
			t.Errorf("Validate after revoke = %v, want ErrTokenExpired", err)
		}
	})
}

func TestRefresh(t *testing.T) {
	ctx := context.Background()
	m := newManager(t, tokenx.Config{})

	first, err := m.Generate(ctx, "u1")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}

	second, err := m.Refresh(ctx, "u1", first.RefreshToken)
	if err != nil {
		t.Fatalf("Refresh: %v", err)
	}
	if second.AccessToken == first.AccessToken {
		t.Error("Refresh 应轮换 access token")
	}
	if err := m.Validate(ctx, "u1", second.AccessToken); err != nil {
		t.Errorf("新 access token 校验失败: %v", err)
	}

	if _, err := m.Refresh(ctx, "u1", first.RefreshToken); !errors.Is(err, tokenx.ErrTokenInvalid) {
		t.Errorf("复用旧 refresh token = %v, want ErrTokenInvalid", err)
	}
}

func TestRevoke_EmptyUserID(t *testing.T) {
	if err := newManager(t, tokenx.Config{}).Revoke(context.Background(), ""); err == nil {
		t.Fatal("expected error for empty userID")
	}
}

// 空 KeyPrefix 也须能正常往返，key 退化为 type:userID。
func TestEmptyKeyPrefix_RoundTrips(t *testing.T) {
	m := newManager(t, tokenx.Config{KeyPrefix: ""})
	pair, err := m.Generate(context.Background(), "u1")
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if err := m.Validate(context.Background(), "u1", pair.AccessToken); err != nil {
		t.Errorf("Validate: %v", err)
	}
}
