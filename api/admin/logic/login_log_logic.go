package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type LoginLogLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewLoginLogLogic(svcCtx *svctx.ServiceContext) *LoginLogLogic {
	return &LoginLogLogic{
		svcCtx: svcCtx,
	}
}

// 批量删除登录日志
func (l *LoginLogLogic) DeleteLoginLog(reqCtx *request.Context, req *types.DeleteLoginLogReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取登录日志列表
func (l *LoginLogLogic) QueryLoginLogList(reqCtx *request.Context, req *types.QueryLoginLogListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}
