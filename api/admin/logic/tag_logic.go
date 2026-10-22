package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type TagLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewTagLogic(svcCtx *svctx.ServiceContext) *TagLogic {
	return &TagLogic{
		svcCtx: svcCtx,
	}
}

// 创建标签
func (l *TagLogic) CreateTag(reqCtx *request.Context, req *types.CreateTagReq) (resp *types.TagVO, err error) {
	// TODO: implement business logic

	return
}

// 删除标签
func (l *TagLogic) DeleteTag(reqCtx *request.Context, req *types.DeleteTagReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取标签列表
func (l *TagLogic) QueryTagList(reqCtx *request.Context, req *types.QueryTagListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 更新标签
func (l *TagLogic) UpdateTag(reqCtx *request.Context, req *types.UpdateTagReq) (resp *types.TagVO, err error) {
	// TODO: implement business logic

	return
}
