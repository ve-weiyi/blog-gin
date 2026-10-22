package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ve-weiyi/blog-gin/infra/biz/bizheader"
	"github.com/ve-weiyi/blog-gin/infra/logz"
)

type contextKey string

const CtxKeyTraceId contextKey = "trace_id"

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		// context注入traceID
		traceId := c.GetHeader(bizheader.HeaderXRequestId)
		if traceId == "" {
			traceId = uuid.New().String()
		}
		c.Set(CtxKeyTraceId, traceId)
		// 同时写入 Request 的 context：gin context 只在 HTTP 层可用，
		// 下沉到 logic 的是 request.Context 内嵌的那个标准 context
		c.Request = c.Request.WithContext(logz.WithTraceID(c.Request.Context(), traceId))

		c.Next()
	}
}
