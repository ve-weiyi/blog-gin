package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/ve-weiyi/ve-blog-gin/infra/biz/bizheader"
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

		c.Next()
	}
}
