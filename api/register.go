package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/vkit/plugins/knife4j"
	"go.opentelemetry.io/otel/trace"

	"github.com/ve-weiyi/blog-gin/api/admin"
	"github.com/ve-weiyi/blog-gin/api/app"
	"github.com/ve-weiyi/blog-gin/docs"
	"github.com/ve-weiyi/blog-gin/infra/middleware"
	"github.com/ve-weiyi/blog-gin/svctx"
)

// 初始化总路由
func RegisterRouters(engine *gin.Engine, svcCtx *svctx.ServiceContext) {
	v1 := engine.Group("/api/v1")
	RegisterPingHandlers(v1, svcCtx)
	RegisterSwaggerHandlers(v1, svcCtx)

	g := engine.Group("")

	// r.Use(middleware.Logger())  // 访问记录，gin.Default()自带了
	// r.Use(middleware.Recovery(true)) // 捕获panic，gin.Default()自带了
	// r.Use(middleware.Limit(svcCtx.RedisEngin)) // 请求限频，一般在nginx层处理。
	g.Use(middleware.Cors())  // 直接放行全部跨域请求
	g.Use(middleware.Trace()) // 打印请求的traceId

	app.RegisterHandlers(g, svcCtx)
	admin.RegisterHandlers(g, svcCtx)
}

// 健康检查
func RegisterPingHandlers(r *gin.RouterGroup, svcCtx *svctx.ServiceContext) {
	now := time.Now()
	// 健康监测
	r.GET("/version", func(c *gin.Context) {
		spanCtx := trace.SpanContextFromContext(c.Request.Context())

		c.JSON(http.StatusOK, gin.H{
			"version":  "v1.0.0",
			"runtime":  now.String(),
			"trace_id": spanCtx.TraceID().String(),
		})
	})
}

func RegisterSwaggerHandlers(r *gin.RouterGroup, svcCtx *svctx.ServiceContext) {
	// Generate Swagger JSON file
	r.GET("/swagger/*any", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		knife4j.NewKnife4jPlugin(docs.SwaggerInfo.ReadDoc()).Handler("/swagger").ServeHTTP(w, r)
	}))
	fmt.Printf("swagger docs at http://localhost:%d/api/v1/swagger/index.html\n", svcCtx.Config.Server.Port)
}
