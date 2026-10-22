package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/logic"
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
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
// @Router /api/v1/user/me/bind_user_email [POST]
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
// @Router /api/v1/user/me/bind_user_phone [POST]
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
// @Router /api/v1/user/me/bind_user_third_party [POST]
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

// 停用账号（进入冷静期）
// @Param data body types.DeactivateAccountReq true "请求参数"
// @Success 200 {object} types.DeactivateAccountResp "返回信息"
// @Router /api/v1/user/me/deactivate_account [POST]
func (h *MeHandler) DeactivateAccount(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeactivateAccountReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.DeactivateAccount(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取用户点赞集合
// @Param data body types.GetUserLikeReq true "请求参数"
// @Success 200 {object} types.GetUserLikeResp "返回信息"
// @Router /api/v1/user/me/get_user_like [GET]
func (h *MeHandler) GetUserLike(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetUserLikeReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.GetUserLike(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取当前用户信息
// @Param data body types.GetUserProfileReq true "请求参数"
// @Success 200 {object} types.GetUserProfileResp "返回信息"
// @Router /api/v1/user/me/get_user_profile [GET]
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

// 恢复账号（冷静期内）
// @Param data body types.ReactivateAccountReq true "请求参数"
// @Success 200 {object} types.ReactivateAccountResp "返回信息"
// @Router /api/v1/user/me/reactivate_account [POST]
func (h *MeHandler) ReactivateAccount(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.ReactivateAccountReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewMeLogic(h.svcCtx)
	resp, err := l.ReactivateAccount(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 解绑第三方平台账号
// @Param data body types.UnbindUserThirdPartyReq true "请求参数"
// @Success 200 {object} types.EmptyResp "返回信息"
// @Router /api/v1/user/me/unbind_user_third_party [POST]
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

// 更新当前用户信息
// @Param data body types.UpdateUserProfileReq true "请求参数"
// @Success 200 {object} types.UpdateUserProfileResp "返回信息"
// @Router /api/v1/user/me/update_user_profile [PUT]
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
