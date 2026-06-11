package app

import (
	"github.com/gin-gonic/gin"
	"github.com/ve-weiyi/ve-blog-gin/api/app/router"
	"github.com/ve-weiyi/ve-blog-gin/svctx"
)

func RegisterHandlers(r *gin.RouterGroup, svcCtx *svctx.ServiceContext) {
	router.NewAlbumRouter(svcCtx).Register(r)
	router.NewArticleRouter(svcCtx).Register(r)
	router.NewAuthRouter(svcCtx).Register(r)
	router.NewCategoryRouter(svcCtx).Register(r)
	router.NewCommentRouter(svcCtx).Register(r)
	router.NewConfigRouter(svcCtx).Register(r)
	router.NewDefaultRouter(svcCtx).Register(r)
	router.NewFriendRouter(svcCtx).Register(r)
	router.NewGuestRouter(svcCtx).Register(r)
	router.NewMeRouter(svcCtx).Register(r)
	router.NewMessageRouter(svcCtx).Register(r)
	router.NewPageRouter(svcCtx).Register(r)
	router.NewTagRouter(svcCtx).Register(r)
	router.NewTalkRouter(svcCtx).Register(r)
	router.NewUploadRouter(svcCtx).Register(r)
	router.NewWebsocketRouter(svcCtx).Register(r)
}
