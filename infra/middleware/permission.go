package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/ve-weiyi/blog-gin/infra/biz/bizcode"
	"github.com/ve-weiyi/blog-gin/infra/biz/bizerr"
	"github.com/ve-weiyi/blog-gin/infra/middleware/permissionx"
	"github.com/ve-weiyi/blog-gin/infra/response"
)

// Permission 权限拦截中间件
func Permission(enforcer permissionx.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := GetUserId(c)
		if uid == "" {
			response.ResponseError(c, bizerr.NewBizError(bizcode.CodeUnauthorized, "request header field 'uid' is missing"))
			c.Abort()
			return
		}

		ok, err := enforcer.Enforce(uid, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			response.ResponseError(c, bizerr.NewBizError(bizcode.CodeNoPermission, err.Error()))
			c.Abort()
			return
		}

		if !ok {
			response.ResponseError(c, bizerr.NewBizError(bizcode.CodeNoPermission, "user does not have permission to access this resource"))
			c.Abort()
			return
		}

		c.Next()
	}
}
