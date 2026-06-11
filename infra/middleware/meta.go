package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/ve-blog-gin/infra/biz/bizheader"
)

// Meta 将请求头中的关键字段注入 Gin context
func Meta() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, key := range bizheader.HeaderFields {
			if v := c.GetHeader(key); v != "" {
				c.Set(key, v)
			}
		}
		c.Next()
	}
}

// GetUserId 从 context 获取用户 ID
func GetUserId(c *gin.Context) string {
	return c.GetString(bizheader.HeaderUid)
}

// GetDeviceId 从 context 获取设备 ID
func GetDeviceId(c *gin.Context) string {
	return c.GetString(bizheader.HeaderXDeviceId)
}
