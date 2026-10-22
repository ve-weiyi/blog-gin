package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
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

// 创建说说
func (l *TalkLogic) CreateTalk(reqCtx *request.Context, req *types.CreateTalkReq) (resp *types.TalkVO, err error) {
	// TODO: implement business logic

	return
}

// 删除说说
func (l *TalkLogic) DeleteTalk(reqCtx *request.Context, req *types.DeleteTalkReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取说说详情
func (l *TalkLogic) GetTalk(reqCtx *request.Context, req *types.GetTalkReq) (resp *types.TalkVO, err error) {
	// TODO: implement business logic

	return
}

// 获取说说列表
func (l *TalkLogic) QueryTalkList(reqCtx *request.Context, req *types.QueryTalkListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 更新说说
func (l *TalkLogic) UpdateTalk(reqCtx *request.Context, req *types.UpdateTalkReq) (resp *types.TalkVO, err error) {
	// TODO: implement business logic

	return
}
