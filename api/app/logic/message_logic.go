package logic

import (
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type MessageLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewMessageLogic(svcCtx *svctx.ServiceContext) *MessageLogic {
	return &MessageLogic{
		svcCtx: svcCtx,
	}
}

// 获取留言列表
func (l *MessageLogic) QueryMessageList(reqCtx *request.Context, req *types.QueryMessageListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 创建留言
func (l *MessageLogic) CreateMessage(reqCtx *request.Context, req *types.CreateMessageReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}
