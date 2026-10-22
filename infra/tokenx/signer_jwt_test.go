package tokenx_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/ve-weiyi/blog-gin/infra/tokenx"
)

func TestJWTSigner_Sign(t *testing.T) {
	signer := tokenx.NewJWTSigner([]byte("my-secret-key"), "test-issuer")
	token, err := signer.Sign(context.Background(), tokenx.TokenClaims{
		UserID:    "user1",
		IssuedAt:  time.Unix(1000, 0),
		ExpiresAt: time.Unix(1900, 0),
	})
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}

	// JWT 由三段以点分隔。
	if parts := len(strings.Split(token, ".")); parts != 3 {
		t.Fatalf("JWT 段数 = %d, want 3: %s", parts, token)
	}
}

func TestJWTSigner_Verify(t *testing.T) {
	signer := tokenx.NewJWTSigner([]byte("secret"), "iss")
	claims := tokenx.TokenClaims{
		UserID:    "user1",
		TokenType: "access",
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Hour),
	}
	token, err := signer.Sign(context.Background(), claims)
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}

	t.Run("valid token", func(t *testing.T) {
		parsed, err := signer.Verify(context.Background(), token)
		if err != nil {
			t.Fatalf("Verify: %v", err)
		}
		if parsed.UserID != "user1" || parsed.TokenType != "access" {
			t.Errorf("claims = %+v", parsed)
		}
	})

	t.Run("expired token", func(t *testing.T) {
		expired, err := signer.Sign(context.Background(), tokenx.TokenClaims{
			UserID:    "u",
			IssuedAt:  time.Now().Add(-2 * time.Hour),
			ExpiresAt: time.Now().Add(-time.Hour),
		})
		if err != nil {
			t.Fatalf("Sign: %v", err)
		}
		if _, err := signer.Verify(context.Background(), expired); err != tokenx.ErrTokenExpired {
			t.Errorf("Verify(expired) = %v, want ErrTokenExpired", err)
		}
	})

	t.Run("malformed token", func(t *testing.T) {
		if _, err := signer.Verify(context.Background(), "not.a.valid.jwt"); err != tokenx.ErrTokenInvalid {
			t.Errorf("Verify(malformed) = %v, want ErrTokenInvalid", err)
		}
	})

	t.Run("wrong key", func(t *testing.T) {
		other := tokenx.NewJWTSigner([]byte("secret2"), "iss")
		if _, err := other.Verify(context.Background(), token); err != tokenx.ErrTokenInvalid {
			t.Errorf("Verify(wrong key) = %v, want ErrTokenInvalid", err)
		}
	})
}
