package logic

import (
	"github.com/ve-weiyi/blog-gin/api/admin/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type ApiLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewApiLogic(svcCtx *svctx.ServiceContext) *ApiLogic {
	return &ApiLogic{
		svcCtx: svcCtx,
	}
}

// 清空接口列表
func (l *ApiLogic) CleanApi(reqCtx *request.Context, req *types.EmptyReq) (resp *types.CleanApiResp, err error) {
	// TODO: implement business logic

	return
}

// 创建接口
func (l *ApiLogic) CreateApi(reqCtx *request.Context, req *types.CreateApiReq) (resp *types.ApiVO, err error) {
	// TODO: implement business logic

	return
}

// 删除接口
func (l *ApiLogic) DeleteApi(reqCtx *request.Context, req *types.DeleteApiReq) (resp *types.BatchResp, err error) {
	// TODO: implement business logic

	return
}

// 获取接口详情
func (l *ApiLogic) GetApi(reqCtx *request.Context, req *types.GetApiReq) (resp *types.ApiVO, err error) {
	// TODO: implement business logic

	return
}

// 获取接口列表
func (l *ApiLogic) QueryApiList(reqCtx *request.Context, req *types.QueryApiListReq) (resp *types.PageResult, err error) {
	// TODO: implement business logic

	return
}

// 同步接口列表
func (l *ApiLogic) SyncApi(reqCtx *request.Context, req *types.EmptyReq) (resp *types.SyncApiResp, err error) {
	// TODO: implement business logic

	return
}

// 更新接口
func (l *ApiLogic) UpdateApi(reqCtx *request.Context, req *types.UpdateApiReq) (resp *types.ApiVO, err error) {
	// TODO: implement business logic

	return
}
