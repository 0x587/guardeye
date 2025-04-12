package logic

import (
	"context"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/samber/lo/parallel"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/ent"
	"github.com/0x587/guardeye/common/rediskey"
	"github.com/0x587/guardeye/report/report"
)

type NodesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNodesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NodesLogic {
	return &NodesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NodesLogic) Nodes(req *types.NodesReq) (resp *types.NodesRsp, err error) {
	agents, err := l.svcCtx.Db.Agent.Query().All(l.ctx)
	if err != nil {
		return nil, err
	}
	nodes := parallel.Map(agents, func(n *ent.Agent, index int) types.NodeInfo {
		res := types.NodeInfo{
			NodeId: n.ClientID.String(),
			Name:   n.Alias,
			Os:     n.Os,
			Arch:   n.OsVersion,
			Macs:   n.Macs,
			Ips:    n.Ips,
			Cpu:    n.CPU,
			Memory: n.Memory,
			Disk:   n.Disk,
			UpTime: n.Uptime,
		}
		if res.Name == "" {
			res.Name = "Unknown Node"
		}
		seen, err := l.svcCtx.Redis.GetCtx(l.ctx, rediskey.SeeAtKey(&report.NodeInfo{ClientId: n.ClientID.String()}))
		if err == nil {
			parseInt, _ := strconv.ParseInt(seen, 10, 64)
			lastSeen := time.Unix(0, parseInt)
			res.LastSeenAt = lastSeen.String()
			if lastSeen.Add(time.Duration(5) * time.Minute).Before(time.Now()) {
				res.UpTime = ""
				res.Status = "offline"
			}
			res.Status = "online"
		} else {
			logx.Error(err)
		}
		latency, err := l.svcCtx.Redis.GetCtx(l.ctx, rediskey.LatencyKey(&report.NodeInfo{ClientId: n.ClientID.String()}))
		logx.Info(latency)
		if err == nil {
			lat, _ := strconv.ParseUint(latency, 10, 64)
			res.Latency = int(lat)
			if lat > 100*1000*1000 { // 100ms
				res.Status = "degraded"
			}
		} else {
			logx.Error(err)
		}
		return res
	})
	nodes = lo.Map(nodes, func(item types.NodeInfo, _ int) types.NodeInfo {
		item.Ips = lo.Uniq(lo.Filter(item.Ips, func(item string, _ int) bool {
			return len(item) > 0
		}))
		//item.Ips = slices.SortFunc(item.Ips, func(a, b string) int {
		//
		//})
		item.Macs = lo.Uniq(lo.Filter(item.Macs, func(item string, _ int) bool {
			return len(item) > 0
		}))
		return item
	})
	slices.SortFunc(nodes, func(a, b types.NodeInfo) int {
		return strings.Compare(a.NodeId, b.NodeId)
	})
	resp = &types.NodesRsp{
		Nodes: nodes,
	}
	return
}
