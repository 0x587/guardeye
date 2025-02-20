package sendtows

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/zeromicro/go-queue/kq"

	"github.com/0x587/guardeye/wsapi/internal/svc"
	"github.com/0x587/guardeye/wsapi/internal/types"
	"github.com/0x587/guardeye/wsapi/internal/ws"
)

func New(ctx context.Context, svcCtx *svc.ServiceContext) kq.ConsumeHandler {
	i := &impl{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
	i.svcCtx.Ws.AddHandler(i.handler)
	return i
}

type impl struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (i *impl) Consume(ctx context.Context, key, val string) error {
	cids, err := i.svcCtx.Redis.Smembers(fmt.Sprintf("wsapi-sub-key-%s", key))
	if err != nil {
		return nil
	}
	for _, cid := range cids {
		i.svcCtx.Ws.Send(uuid.MustParse(cid), []byte(val))
	}
	return nil
}

func (i *impl) handler(ws ws.IF, cid uuid.UUID, msg []byte) {
	m := &types.WsMsgBase{}
	if err := json.Unmarshal(msg, m); err != nil {
		return
	}
	switch m.Cmd {
	case "Sub":
		req := &types.WsSubReq{}
		if err := json.Unmarshal(msg, req); err != nil {
			return
		}
		if req.Action == "Sub" {
			if _, err := i.svcCtx.Redis.Sadd(fmt.Sprintf("wsapi-sub-key-%s", req.SubKey), cid.String()); err != nil {
				return
			}
		}
		if req.Action == "Unsub" {
			if _, err := i.svcCtx.Redis.Srem(fmt.Sprintf("wsapi-sub-key-%s", req.SubKey), cid.String()); err != nil {
				return
			}
		}
	}
}
