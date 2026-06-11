package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type MeHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewMeHandler(svcCtx *svctx.ServiceContext) *MeHandler {
	return &MeHandler{
		svcCtx: svcCtx,
	}
}

// 绑定邮箱
// @Param data body types.BindUserEmailReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/user/me/bind_user_email [POST]
func (h *MeHandler) BindUserEmail(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.BindUserEmailReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.BindUserEmail(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 绑定手机号
// @Param data body types.BindUserPhoneReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/user/me/bind_user_phone [POST]
func (h *MeHandler) BindUserPhone(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.BindUserPhoneReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.BindUserPhone(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 绑定第三方平台账号
// @Param data body types.BindUserThirdPartyReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/user/me/bind_user_third_party [POST]
func (h *MeHandler) BindUserThirdParty(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.BindUserThirdPartyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.BindUserThirdParty(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取用户接口权限
// @Param data body types.EmptyReq true "请求参数"
// @Success 200 {object} types.GetUserApisResp "返回信息"
// @Router /admin-api/v1/user/me/get_user_apis [GET]
func (h *MeHandler) GetUserApis(c *gin.Context) {
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

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.GetUserApis(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取用户菜单权限
// @Param data body types.EmptyReq true "请求参数"
// @Success 200 {object} types.GetUserMenusResp "返回信息"
// @Router /admin-api/v1/user/me/get_user_menus [GET]
func (h *MeHandler) GetUserMenus(c *gin.Context) {
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

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.GetUserMenus(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取当前用户信息
// @Param data body types.GetUserProfileReq true "请求参数"
// @Success 200 {object} types.UserProfile "返回信息"
// @Router /admin-api/v1/user/me/get_user_profile [GET]
func (h *MeHandler) GetUserProfile(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetUserProfileReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.GetUserProfile(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取用户角色
// @Param data body types.EmptyReq true "请求参数"
// @Success 200 {object} types.GetUserRolesResp "返回信息"
// @Router /admin-api/v1/user/me/get_user_roles [GET]
func (h *MeHandler) GetUserRoles(c *gin.Context) {
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

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.GetUserRoles(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 查询用户登录历史
// @Param data body types.QueryUserLoginHistoryReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/user/me/query_user_login_history [POST]
func (h *MeHandler) QueryUserLoginHistory(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryUserLoginHistoryReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.QueryUserLoginHistory(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 解绑第三方平台账号
// @Param data body types.UnbindUserThirdPartyReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/user/me/unbind_user_third_party [POST]
func (h *MeHandler) UnbindUserThirdParty(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UnbindUserThirdPartyReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.UnbindUserThirdParty(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 修改用户头像
// @Param data body types.UpdateUserAvatarReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/user/me/update_user_avatar [PUT]
func (h *MeHandler) UpdateUserAvatar(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateUserAvatarReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.UpdateUserAvatar(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 修改用户密码
// @Param data body types.UpdateUserPasswordReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/user/me/update_user_password [PUT]
func (h *MeHandler) UpdateUserPassword(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateUserPasswordReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.UpdateUserPassword(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新当前用户信息
// @Param data body types.UpdateUserProfileReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /admin-api/v1/user/me/update_user_profile [PUT]
func (h *MeHandler) UpdateUserProfile(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateUserProfileReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.UpdateUserProfile(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
