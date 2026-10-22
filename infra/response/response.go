package response

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"

	"github.com/ve-weiyi/blog-gin/infra/biz/bizcode"
	"github.com/ve-weiyi/blog-gin/infra/biz/bizerr"
)

type Body struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(c *gin.Context, code int64, msg string, data interface{}) {
	c.JSON(int(bizcode.CodeSuccess), Body{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func ResponseOk(c *gin.Context, data interface{}) {
	Response(c, bizcode.CodeSuccess, "ok", data)
}

func ResponseError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	// 1. 返回错误信息多语言
	// 2. 忽略详细的内部错误信息
	// 3. 处理不同的业务逻辑
	var bizErr *bizerr.BizError
	var mysqlErr *mysql.MySQLError

	switch {
	case errors.As(err, &bizErr):
		Response(c, int64(bizErr.Code), bizErr.Error(), nil)
	case errors.As(err, &mysqlErr):
		Response(c, bizcode.CodeServerError, SqlErrorI18n(mysqlErr), nil)
	default:
		Response(c, bizcode.CodeServerError, err.Error(), nil)
	}
}

// Error 通过业务码快速返回错误响应
func Error(c *gin.Context, code int) {
	ResponseError(c, bizerr.NewBizError(code, bizcode.Message(code)))
}
