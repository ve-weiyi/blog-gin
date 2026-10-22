package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/blog-gin/infra/biz/bizheader"
	"github.com/ve-weiyi/blog-gin/infra/logz"
)

func TestTrace_SeedsTraceID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name      string
		header    string
		wantExact bool
	}{
		{name: "沿用请求头 x-request-id", header: "upstream-trace-1", wantExact: true},
		{name: "无请求头时生成 uuid", header: "", wantExact: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			engine := gin.New()
			engine.Use(Trace())

			var fromCtx string
			var fromGin any
			engine.GET("/ping", func(c *gin.Context) {
				fromCtx = logz.TraceIDFrom(c.Request.Context())
				fromGin, _ = c.Get(CtxKeyTraceId)
				c.Status(http.StatusOK)
			})

			req := httptest.NewRequest(http.MethodGet, "/ping", nil)
			if tt.header != "" {
				req.Header.Set(bizheader.HeaderXRequestId, tt.header)
			}
			engine.ServeHTTP(httptest.NewRecorder(), req)

			if tt.wantExact {
				if fromCtx != tt.header {
					t.Fatalf("expected traceId %q in request context, got %q", tt.header, fromCtx)
				}
				return
			}
			if fromCtx == "" {
				t.Fatal("expected generated traceId in request context, got empty")
			}
			// gin context 与 request context 必须是同一个值，否则日志拿到的与响应头写入的不是同一个 traceId
			if got, _ := fromGin.(string); got != fromCtx {
				t.Fatalf("gin context traceId %q != request context traceId %q", got, fromCtx)
			}
		})
	}
}

// 无请求头时生成的 traceId 须两两不同。
func TestTrace_GeneratedTraceIDsAreUnique(t *testing.T) {
	gin.SetMode(gin.TestMode)

	engine := gin.New()
	engine.Use(Trace())

	var ids []string
	engine.GET("/ping", func(c *gin.Context) {
		ids = append(ids, logz.TraceIDFrom(c.Request.Context()))
		c.Status(http.StatusOK)
	})

	for i := 0; i < 5; i++ {
		engine.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/ping", nil))
	}

	seen := make(map[string]bool, len(ids))
	for _, id := range ids {
		if id == "" {
			t.Fatal("generated traceId must not be empty")
		}
		if seen[id] {
			t.Fatalf("duplicated traceId %q", id)
		}
		seen[id] = true
	}
}
