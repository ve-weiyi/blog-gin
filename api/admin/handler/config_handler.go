package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type ConfigHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewConfigHandler(svcCtx *svctx.ServiceContext) *ConfigHandler {
	return &ConfigHandler{
		svcCtx: svcCtx,
	}
}

// 获取关于我的信息
// @Param data body types.EmptyReq true "请求参数"
// @Success 200 {object} types.AboutMeVO "返回信息"
// @Router /admin-api/v1/admin/get_about_me [GET]
func (h *ConfigHandler) GetAboutMe(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.EmptyReq
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

// 获取服务器信息
// @Param data body types.EmptyReq true "请求参数"
// @Success 200 {object} types.Server "返回信息"
// @Router /admin-api/v1/admin/get_system_state [GET]
func (h *ConfigHandler) GetSystemState(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewConfigLogic(h.svcCtx)
	resp, err := l.GetSystemState(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取网站配置
// @Param data body types.EmptyReq true "请求参数"
// @Success 200 {object} types.WebsiteConfigVO "返回信息"
// @Router /admin-api/v1/admin/get_website_config [GET]
func (h *ConfigHandler) GetWebsiteConfig(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.EmptyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewConfigLogic(h.svcCtx)
	resp, err := l.GetWebsiteConfig(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新关于我的信息
// @Param data body types.AboutMeVO true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/admin/update_about_me [PUT]
func (h *ConfigHandler) UpdateAboutMe(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.AboutMeVO
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewConfigLogic(h.svcCtx)
	resp, err := l.UpdateAboutMe(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新网站配置
// @Param data body types.WebsiteConfigVO true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/admin/update_website_config [PUT]
func (h *ConfigHandler) UpdateWebsiteConfig(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.WebsiteConfigVO
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewConfigLogic(h.svcCtx)
	resp, err := l.UpdateWebsiteConfig(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
