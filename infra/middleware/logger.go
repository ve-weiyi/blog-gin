package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 接收gin框架默认的日志
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path      // 请求路径 eg: /test
		query := c.Request.URL.RawQuery //query类型的请求参数：?name=1&password=2
		// 挂起当前中间件，执行下一个中间件
		c.Next()
		// 视图函数执行完成，统计时间，记录日志
		cost := time.Since(start)

		log.Printf("[GIN] status=%d method=%s path=%s query=%s ip=%s user-agent=%s errors=%s cost=%v",
			c.Writer.Status(),
			c.Request.Method,
			path,
			query,
			c.ClientIP(),
			c.Request.UserAgent(),
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
			cost,
		)
	}
}
