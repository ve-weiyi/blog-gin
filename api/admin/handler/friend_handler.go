package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/logic"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/types"
	"github.com/ve-weiyi/ve-blog-gin/infra/request"
	"github.com/ve-weiyi/ve-blog-gin/infra/response"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

type FriendHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewFriendHandler(svcCtx *svctx.ServiceContext) *FriendHandler {
	return &FriendHandler{
		svcCtx: svcCtx,
	}
}

// 创建友链
// @Param data body types.CreateFriendReq true "请求参数"
// @Success 200 {object} types.FriendVO "返回信息"
// @Router /admin-api/v1/friend/create_friend [POST]
func (h *FriendHandler) CreateFriend(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.CreateFriendReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewFriendLogic(h.svcCtx)
	resp, err := l.CreateFriend(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 批量删除友链
// @Param data body types.DeleteFriendReq true "请求参数"
// @Success 200 {object} types.BatchResp "返回信息"
// @Router /admin-api/v1/friend/delete_friend [DELETE]
func (h *FriendHandler) DeleteFriend(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.DeleteFriendReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewFriendLogic(h.svcCtx)
	resp, err := l.DeleteFriend(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 获取友链列表
// @Param data body types.QueryFriendListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /admin-api/v1/friend/query_friend_list [POST]
func (h *FriendHandler) QueryFriendList(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.QueryFriendListReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewFriendLogic(h.svcCtx)
	resp, err := l.QueryFriendList(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}

// 更新友链
// @Param data body types.UpdateFriendReq true "请求参数"
// @Success 200 {object} types.FriendVO "返回信息"
// @Router /admin-api/v1/friend/update_friend [PUT]
func (h *FriendHandler) UpdateFriend(c *gin.Context) {
	reqCtx, err := request.ParseRequestContext(c)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	var req *types.UpdateFriendReq
	err = request.ShouldBind(c, &req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}

	l := logic.NewFriendLogic(h.svcCtx)
	resp, err := l.UpdateFriend(reqCtx, req)
	if err != nil {
		response.ResponseError(c, err)
		return
	}
	response.ResponseOk(c, resp)
}
