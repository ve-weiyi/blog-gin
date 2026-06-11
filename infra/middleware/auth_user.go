package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/vkit/adapter/storex/tokenstore"

	"github.com/ve-weiyi/ve-blog-gin/infra/biz/bizcode"
	"github.com/ve-weiyi/ve-blog-gin/infra/biz/bizheader"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
)

const bearerPrefix = "Bearer "

// UserAuth 前台用户 JWT 鉴权
func UserAuth(tokenStore tokenstore.TokenStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetHeader(bizheader.HeaderUid)
		token := c.GetHeader(bizheader.HeaderAuthorization)
		if token == "" {
			token = c.Query(bizheader.HeaderToken)
		}
		if uid == "" || token == "" {
			response.Error(c, bizcode.CodeUnauthorized)
			c.Abort()
			return
		}

		if len(token) > len(bearerPrefix) && token[:len(bearerPrefix)] == bearerPrefix {
			token = token[len(bearerPrefix):]
		}

		if err := tokenStore.ValidateToken(uid, token); err != nil {
			response.Error(c, bizcode.CodeUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
