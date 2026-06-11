package logic

import (
	"github.com/ve-weiyi/ve-blog-gin/api/app/types"
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

// 获取博客前台首页信息
func (l *ConfigLogic) GetBlogHomeInfo(reqCtx *request.Context, req *types.GetBlogHomeInfoReq) (resp *types.GetBlogHomeInfoResp, err error) {
	// TODO: implement business logic

	return
}

// 获取关于我的信息
func (l *ConfigLogic) GetAboutMe(reqCtx *request.Context, req *types.GetAboutMeReq) (resp *types.GetAboutMeResp, err error) {
	// TODO: implement business logic

	return
}
