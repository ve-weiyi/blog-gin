package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type ConfigLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewConfigLogic(svcCtx *svctx.ServiceContext) *ConfigLogic {
	return &ConfigLogic{
		svcCtx: svcCtx,
	}
}

// 获取关于我的信息
func (l *ConfigLogic) GetAboutMe(reqCtx *request.Context, req *types.EmptyReq) (resp *types.AboutMeVO, err error) {
	// TODO: implement business logic

	return
}

// 获取服务器信息
func (l *ConfigLogic) GetSystemState(reqCtx *request.Context, req *types.EmptyReq) (resp *types.Server, err error) {
	// TODO: implement business logic

	return
}

// 获取网站配置
func (l *ConfigLogic) GetWebsiteConfig(reqCtx *request.Context, req *types.EmptyReq) (resp *types.WebsiteConfigVO, err error) {
	// TODO: implement business logic

	return
}

// 更新关于我的信息
func (l *ConfigLogic) UpdateAboutMe(reqCtx *request.Context, req *types.AboutMeVO) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}

// 更新网站配置
func (l *ConfigLogic) UpdateWebsiteConfig(reqCtx *request.Context, req *types.WebsiteConfigVO) (resp *types.EmptyResp, err error) {
	// TODO: implement business logic

	return
}
