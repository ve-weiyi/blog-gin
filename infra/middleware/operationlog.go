package middleware

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// responseRecorder 记录响应体和状态码
type responseRecorder struct {
	gin.ResponseWriter
	body       bytes.Buffer
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

func (rr *responseRecorder) Write(b []byte) (int, error) {
	rr.body.Write(b)
	return rr.ResponseWriter.Write(b)
}

// OperationLog 操作日志记录中间件
func OperationLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// 读取请求体
		var reqBody string
		if c.Request.Body != nil {
			body, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			reqBody = string(body)
		}

		// 替换 ResponseWriter 以记录响应
		rec := &responseRecorder{ResponseWriter: c.Writer, statusCode: 200}
		c.Writer = rec

		c.Next()

		cost := time.Since(start)

		respBody := rec.body.String()
		maxLen := 1000
		if len(reqBody) > maxLen {
			reqBody = reqBody[:maxLen]
		}
		if len(respBody) > maxLen {
			respBody = respBody[:maxLen]
		}

		log.Printf("[OperationLog] uid=%s method=%s path=%s status=%d cost=%v req=%s resp=%s",
			GetUserId(c),
			c.Request.Method,
			c.Request.URL.Path,
			rec.statusCode,
			cost,
			reqBody,
			respBody,
		)
	}
}
