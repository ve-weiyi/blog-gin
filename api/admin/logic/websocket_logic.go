package logic

import (
	"github.com/ve-weiyi/blog-gin/infra/request"
	"github.com/ve-weiyi/blog-gin/svctx"
)

type WebsocketLogic struct {
	svcCtx *svctx.ServiceContext
}

func NewWebsocketLogic(svcCtx *svctx.ServiceContext) *WebsocketLogic {
	return &WebsocketLogic{
		svcCtx: svcCtx,
	}
}

// WebSocket消息
func (l *WebsocketLogic) Websocket(reqCtx *request.Context) (err error) {
	// TODO: implement business logic

	return
}
