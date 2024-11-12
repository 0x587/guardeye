package logtodb

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/0x587/guardeye/common/async"
	"github.com/0x587/guardeye/common/model"
	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
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
	logx.Infof("consume key: %s, val: %s\n", key, val)
	return async.GoAndWait(ctx,
		func() error {
			// 将原始日志落库
			return i.svcCtx.RawLogDBClient.Insert(ctx, &model.RawLog{
				ClientID: d.GetNodeInfo().GetClientId(),
				Message:  d.GetMessage(),
				Provider: d.GetProvider(),
			})
		},
		func() error {
			// 追踪上报的node info变化
			nodeInfo := d.GetNodeInfo()
			if nodeInfo == nil {
				return nil
			}
			node, err := i.svcCtx.NodeDBClient.FindOneWithClientID(ctx, nodeInfo.GetClientId())
			if err != nil && !errors.Is(err, model.ErrNotFound) {
				return err
			}
			desc := nodeInfo.GetNodeDescription()
			if desc == nil {
				return nil
			}
			if node != nil && proto.Equal(node.Description, desc) {
				return nil
			}
			return i.svcCtx.NodeDBClient.Insert(ctx, &model.Node{
				ClientID:    nodeInfo.GetClientId(),
				Description: desc,
				Ips:         desc.GetIps(),
			})
		},
	)
}
