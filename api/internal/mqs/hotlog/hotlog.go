package hotlog

import (
	"context"
	"encoding/json"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/report/report"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/zeromicro/go-queue/kq"
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
	i.svcCtx.LogDispatcher.Handle(
		lo.Must(uuid.Parse(d.GetNodeInfo().GetClientId())),
		d.GetLog(),
	)
	return nil
}
