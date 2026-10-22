package logz

import (
	"context"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestWithTraceID_RoundTrip(t *testing.T) {
	tests := []struct {
		name    string
		traceID string
	}{
		{name: "常见 traceId", traceID: "6f1e0b8a-2c3d-4e5f-8a9b-0c1d2e3f4a5b"},
		{name: "空 traceId", traceID: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := WithTraceID(context.Background(), tt.traceID)
			if got := TraceIDFrom(ctx); got != tt.traceID {
				t.Fatalf("expected %q, got %q", tt.traceID, got)
			}
		})
	}
}

func TestTraceIDFrom_MissingKey_ReturnsEmpty(t *testing.T) {
	if got := TraceIDFrom(context.Background()); got != "" {
		t.Fatalf("expected empty string, got %q", got)
	}
}

// WithTrace 必须把 traceId 写进日志字段，这是本 helper 存在的唯一理由。
func TestWithTrace_AttachesTraceIDField(t *testing.T) {
	core, logs := observer.New(zapcore.InfoLevel)
	logger := zap.New(core)

	const traceID = "trace-abc-123"
	ctx := WithTraceID(context.Background(), traceID)

	WithTrace(ctx, logger).Info("hello")

	entries := logs.All()
	if len(entries) != 1 {
		t.Fatalf("expected 1 log entry, got %d", len(entries))
	}
	got, ok := entries[0].ContextMap()[TraceIDKey]
	if !ok {
		t.Fatalf("expected field %q in log entry, got %v", TraceIDKey, entries[0].ContextMap())
	}
	if got != traceID {
		t.Fatalf("expected %q, got %v", traceID, got)
	}
}

// ctx 中没有 traceId 时不得附加字段，也不得改写 logger 本身。
func TestWithTrace_NoTraceID_ReturnsSameLogger(t *testing.T) {
	core, logs := observer.New(zapcore.InfoLevel)
	logger := zap.New(core)

	got := WithTrace(context.Background(), logger)
	if got != logger {
		t.Fatal("expected the same logger instance when ctx carries no traceId")
	}

	got.Info("hello")

	entries := logs.All()
	if len(entries) != 1 {
		t.Fatalf("expected 1 log entry, got %d", len(entries))
	}
	if _, ok := entries[0].ContextMap()[TraceIDKey]; ok {
		t.Fatalf("expected no %q field, got %v", TraceIDKey, entries[0].ContextMap())
	}
}

// 写入 ctx 的 traceId 不跨 goroutine 泄漏：派生 ctx 各自独立。
func TestWithTraceID_DerivedContextsAreIndependent(t *testing.T) {
	base := WithTraceID(context.Background(), "base-trace")

	child := WithTraceID(base, "child-trace")
	if got := TraceIDFrom(child); got != "child-trace" {
		t.Fatalf("expected child-trace, got %q", got)
	}
	if got := TraceIDFrom(base); got != "base-trace" {
		t.Fatalf("expected base context unchanged, got %q", got)
	}
}
