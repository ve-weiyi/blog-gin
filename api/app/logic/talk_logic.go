package logic

import (
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type TalkLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewTalkLogic(svcCtx *svctx.ServiceContext) *TalkLogic {
	return &TalkLogic{
		svcCtx: svcCtx,
	}
}

// 获取说说详情
func (l *TalkLogic) GetTalk(reqCtx *request.Context, req *types.GetTalkReq) (resp *types.Talk, err error) {
	// TODO: implement business logic

	return
}

// 获取说说列表
func (l *TalkLogic) QueryTalkList(reqCtx *request.Context, req *types.QueryTalkListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 点赞说说
func (l *TalkLogic) LikeTalk(reqCtx *request.Context, req *types.LikeTalkReq) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}
