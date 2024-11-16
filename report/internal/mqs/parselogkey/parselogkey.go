package parselogkey

import (
	"context"
	"encoding/json"

	"github.com/0x587/guardeye/common/tokv"
	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"
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
	var keys []string
	switch d.GetLog().GetType() {
	case report.LogType_TEXT:
		keys = []string{"TEXT"}
	case report.LogType_YAML:
		keys = lo.Keys(tokv.YamlToKv(d.GetLog().GetMessage()))
	case report.LogType_JSON:
		keys = lo.Keys(tokv.JsonToKv(d.GetLog().GetMessage()))
	}
	return i.svcCtx.DataKeyRedisClient.SetKey(
		ctx,
		d.GetNodeInfo(),
		d.GetLog().GetProvider(),
		keys,
	)
}
