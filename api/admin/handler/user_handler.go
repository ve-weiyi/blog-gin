package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type UserHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewUserHandler(svcCtx *svctx.ServiceContext) *UserHandler {
	return &UserHandler{
		svcCtx: svcCtx,
	}
}

// 获取在线用户列表
// @Param data body types.EmptyReq true "请求参数"
// @Success 200 {object} types.OnlineUserListResp "返回信息"
// @Router /admin-api/v1/user/get_online_users [GET]
func (h *UserHandler) GetOnlineUsers(c *gin.Context) {
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

	l := logic.NewUserLogic(h.svcCtx)
	resp, err := l.GetOnlineUsers(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取用户详情
// @Param data body types.GetUserDetailReq true "请求参数"
// @Success 200 {object} types.GetUserDetailResp "返回信息"
// @Router /admin-api/v1/user/get_user [GET]
func (h *UserHandler) GetUserDetail(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.GetUserDetailReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewUserLogic(h.svcCtx)
	resp, err := l.GetUserDetail(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取用户列表
// @Param data body types.QueryUserListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/user/query_user_list [POST]
func (h *UserHandler) QueryUserList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryUserListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewUserLogic(h.svcCtx)
	resp, err := l.QueryUserList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 重置用户密码
// @Param data body types.ResetUserPasswordReq true "请求参数"
// @Success 200 {object} types.ResetUserPasswordResp "返回信息"
// @Router /admin-api/v1/user/reset_user_password [POST]
func (h *UserHandler) ResetUserPassword(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.ResetUserPasswordReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewUserLogic(h.svcCtx)
	resp, err := l.ResetUserPassword(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新用户角色
// @Param data body types.UpdateUserRolesReq true "请求参数"
// @Success 200 {object} types.UpdateUserRolesResp "返回信息"
// @Router /admin-api/v1/user/update_user_roles [PUT]
func (h *UserHandler) UpdateUserRoles(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateUserRolesReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewUserLogic(h.svcCtx)
	resp, err := l.UpdateUserRoles(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新用户状态
// @Param data body types.UpdateUserStatusReq true "请求参数"
// @Success 200 {object} types.UpdateUserStatusResp "返回信息"
// @Router /admin-api/v1/user/update_user_status [PUT]
func (h *UserHandler) UpdateUserStatus(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateUserStatusReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewUserLogic(h.svcCtx)
	resp, err := l.UpdateUserStatus(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
