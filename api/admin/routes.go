package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/admin/router"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

func RegisterHandlers(r *gin.RouterGroup, svcCtx *svctx.ServiceContext) {
	router.NewAlbumRouter(svcCtx).Register(r)
	router.NewApiRouter(svcCtx).Register(r)
	router.NewArticleRouter(svcCtx).Register(r)
	router.NewAuthRouter(svcCtx).Register(r)
	router.NewCategoryRouter(svcCtx).Register(r)
	router.NewCommentRouter(svcCtx).Register(r)
	router.NewConfigRouter(svcCtx).Register(r)
	router.NewDefaultRouter(svcCtx).Register(r)
	router.NewFriendRouter(svcCtx).Register(r)
	router.NewGuestRouter(svcCtx).Register(r)
	router.NewLoginLogRouter(svcCtx).Register(r)
	router.NewMeRouter(svcCtx).Register(r)
	router.NewMenuRouter(svcCtx).Register(r)
	router.NewMessageRouter(svcCtx).Register(r)
	router.NewNotifyMessageRouter(svcCtx).Register(r)
	router.NewNotifyRecordRouter(svcCtx).Register(r)
	router.NewNotifyTemplateRouter(svcCtx).Register(r)
	router.NewOperationLogRouter(svcCtx).Register(r)
	router.NewPageRouter(svcCtx).Register(r)
	router.NewPhotoRouter(svcCtx).Register(r)
	router.NewRoleRouter(svcCtx).Register(r)
	router.NewStatsRouter(svcCtx).Register(r)
	router.NewTagRouter(svcCtx).Register(r)
	router.NewTalkRouter(svcCtx).Register(r)
	router.NewUploadRouter(svcCtx).Register(r)
	router.NewUploadLogRouter(svcCtx).Register(r)
	router.NewUserRouter(svcCtx).Register(r)
	router.NewVisitLogRouter(svcCtx).Register(r)
	router.NewWebsocketRouter(svcCtx).Register(r)
}
