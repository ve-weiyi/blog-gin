package logz

import (
	"context"

	"go.uber.org/zap"
)

// TraceIDKey traceId 在日志中的字段名
const TraceIDKey = "traceId"

// traceIDCtxKey traceId 在 context 中的键。
// 用未导出类型作键，避免与其它包写入的同名键互相覆盖。
type traceIDCtxKey struct{}

// WithTraceID 把 traceId 写入 ctx，由请求入口的中间件调用。
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDCtxKey{}, traceID)
}

// TraceIDFrom 从 ctx 取 traceId，不存在时返回空字符串。
func TraceIDFrom(ctx context.Context) string {
	traceID, _ := ctx.Value(traceIDCtxKey{}).(string)
	return traceID
}

// WithTrace 返回携带 traceId 字段的 logger。
// ctx 中无 traceId 时原样返回 logger，调用方无需自行分支。
func WithTrace(ctx context.Context, logger *zap.Logger) *zap.Logger {
	traceID := TraceIDFrom(ctx)
	if traceID == "" {
		return logger
	}
	return logger.With(zap.String(TraceIDKey, traceID))
}
