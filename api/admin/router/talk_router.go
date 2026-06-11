package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/handler"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
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
		g := rg.Group("/admin-api/v1")
		g.Use(r.svcCtx.AdminAuth)
		g.Use(r.svcCtx.Permission)
		g.Use(r.svcCtx.OperationLog)

		h := handler.NewTalkHandler(r.svcCtx)
		// 创建说说
		g.POST("/talk/create_talk", h.CreateTalk)
		// 删除说说
		g.DELETE("/talk/delete_talk", h.DeleteTalk)
		// 获取说说详情
		g.GET("/talk/get_talk", h.GetTalk)
		// 获取说说列表
		g.POST("/talk/query_talk_list", h.QueryTalkList)
		// 更新说说
		g.PUT("/talk/update_talk", h.UpdateTalk)
	}
}
