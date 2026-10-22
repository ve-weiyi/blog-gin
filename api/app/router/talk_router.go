package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/blog-gin/api/app/handler"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type TalkRouter struct {
	svcCtx *svctx.ServiceContext
}

func NewTalkRouter(svcCtx *svctx.ServiceContext) *TalkRouter {
	return &TalkRouter{
		svcCtx: svcCtx,
	}
}

func (r *TalkRouter) Register(rg *gin.RouterGroup) {
	{
		g := rg.Group("/api/v1")

		h := handler.NewTalkHandler(r.svcCtx)
		// 获取说说详情
		g.GET("/talk/get_talk", h.GetTalk)
		// 获取说说列表
		g.POST("/talk/query_talk_list", h.QueryTalkList)
	}
	{
		g := rg.Group("/api/v1")
		g.Use(r.svcCtx.UserAuth)

		h := handler.NewTalkHandler(r.svcCtx)
		// 点赞说说
		g.POST("/talk/like_talk", h.LikeTalk)
	}
}
