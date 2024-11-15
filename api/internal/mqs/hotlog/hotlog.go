package hotlog

import (
	"context"
	"encoding/json"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/report/report"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
)

func New(ctx context.Context, svcCtx *svc.ServiceContext) kq.ConsumeHandler {
	return &impl{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type impl struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (i *impl) Consume(ctx context.Context, key, val string) error {
	d := &report.MQLog{}
	if err := json.Unmarshal([]byte(val), d); err != nil {
		return err
	}
	logx.Infof("consume key: %s, val: %s\n", key, val)
	ws := i.svcCtx.BoardCaseWs[d.GetNodeInfo().GetClientId()]
	if ws == nil {
		return nil
	}
	return ws.Broadcast(d.GetMessage())
}
