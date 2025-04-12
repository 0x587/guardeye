package link

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/link/linkclient"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type LinkListenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLinkListenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LinkListenLogic {
	return &LinkListenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LinkListenLogic) LinkListen(req *types.LinkListenReq, w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	linkConn, err := l.svcCtx.LinkCli.AgentListen(l.ctx, &linkclient.AgentListenReq{
		Cid: req.Cid,
	})
	if err != nil {
		return err
	}
	for {
		recv, err := linkConn.Recv()
		if err != nil {
			return err
		}
		marshal, err := json.Marshal(recv)
		if err != nil {
			return err
		}
		if err := conn.WriteMessage(websocket.TextMessage, marshal); err != nil {
			return err
		}
	}
}
