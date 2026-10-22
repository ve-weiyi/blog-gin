package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/logic"
	"github.com/ve-weiyi/blog-gin/api/app/types"
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/infra/response"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type FriendHandler struct {
	svcCtx *svctx.ServiceContext
}

func NewFriendHandler(svcCtx *svctx.ServiceContext) *FriendHandler {
	return &FriendHandler{
		svcCtx: svcCtx,
	}
}

// 获取友链列表
// @Param data body types.QueryFriendListReq true "请求参数"
// @Success 200 {object} types.PageResult "返回信息"
// @Router /api/v1/friend/query_friend_list [POST]
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
