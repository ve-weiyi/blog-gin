package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/logic"
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type ConfigHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewConfigHandler(svcCtx *svctx.ServiceContext) *ConfigHandler {
	return &ConfigHandler{
		svcCtx: svcCtx,
	}
}

// 获取博客前台首页信息
// @Param data body types.GetBlogHomeInfoReq true "请求参数"
// @Success 200 {object} types.GetBlogHomeInfoResp "返回信息"
// @Router /api/v1/blog [GET]
func (h *ConfigHandler) GetBlogHomeInfo(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetBlogHomeInfoReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewConfigLogic(h.svcCtx)
	resp, err := l.GetBlogHomeInfo(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取关于我的信息
// @Param data body types.GetAboutMeReq true "请求参数"
// @Success 200 {object} types.GetAboutMeResp "返回信息"
// @Router /api/v1/blog/get_about_me [GET]
func (h *ConfigHandler) GetAboutMe(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetAboutMeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewConfigLogic(h.svcCtx)
	resp, err := l.GetAboutMe(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
