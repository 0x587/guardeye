package logtodb

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/0x587/guardeye/common/async"
	"github.com/0x587/guardeye/common/model"
	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"
	"github.com/google/uuid"
	"github.com/zeromicro/go-queue/kq"
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
	d := &report.MQLog{}
	if err := json.Unmarshal([]byte(val), d); err != nil {
		return err
	}
	return async.GoAndWait(ctx,
		func() error {
			// 将原始日志落库
			cid, err := uuid.Parse(d.GetNodeInfo().GetClientId())
			if err != nil {
				return err
			}
			_, err = i.svcCtx.RawLogDBClient.Insert(ctx, &model.Rawlog{
				Id:           uuid.New(),
				ClientId:     cid,
				Message:      d.GetLog().GetMessage(),
				ProviderType: d.GetLog().GetProvider().GetType(),
				ProviderArgs: d.GetLog().GetProvider().GetArgs(),
			})
			return err
		},
		func() error {
			// 追踪上报的node info变化
			nodeInfo := d.GetNodeInfo()
			if nodeInfo == nil {
				return nil
			}
			cid, err := uuid.Parse(nodeInfo.GetClientId())
			if err != nil {
				return err
			}
			node, err := i.svcCtx.NodeDBClient.FindOneWithClientID(ctx, cid)
			if err != nil && !errors.Is(err, model.ErrNotFound) {
				return err
			}
			desc := nodeInfo.GetNodeDescription()
			if desc == nil {
				return nil
			}
			if node == nil {
				return nil
			}
			if proto.Equal(&report.NodeDescription{
				Ips:       node.Ips,
				Macs:      node.Macs,
				Os:        node.Os,
				OsVersion: node.OsVersion,
				Hostname:  node.Hostname,
			}, desc) {
				return nil
			}
			_, err = i.svcCtx.NodeDBClient.Insert(ctx, &model.Node{
				Id:        uuid.New(),
				ClientId:  cid,
				Ips:       desc.GetIps(),
				Macs:      desc.GetMacs(),
				Os:        desc.GetOs(),
				OsVersion: desc.GetOsVersion(),
				Hostname:  desc.GetHostname(),
			})
			return err
		},
	)
}
