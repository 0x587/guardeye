package logtodb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/0x587/guardeye/common/model"
	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"
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
	d := &report.LogReportReq{}
	if err := json.Unmarshal([]byte(val), d); err != nil {
		return err
	}
	fmt.Printf("consume key: %s, val: %s\n", key, val)
	err := i.svcCtx.RawLogDBClient.Insert(ctx, &model.RawLog{
		ClientID: d.GetNodeInfo().GetClientId(),
		Message:  d.GetMessage(),
		Provider: d.GetProvider(),
	})
	if err != nil {
		return err
	}
	// TODO check node info
	return nil
}
