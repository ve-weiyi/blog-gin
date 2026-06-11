package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type MessageLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewMessageLogic(svcCtx *svctx.ServiceContext) *MessageLogic {
	return &MessageLogic{
		svcCtx: svcCtx,
	}
}

// 批量删除留言
func (l *MessageLogic) DeleteMessage(reqCtx *request.Context, req *types.DeleteMessageReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取留言列表
func (l *MessageLogic) QueryMessageList(reqCtx *request.Context, req *types.QueryMessageListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 批量更新留言状态
func (l *MessageLogic) UpdateMessageStatus(reqCtx *request.Context, req *types.UpdateMessageStatusReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}
