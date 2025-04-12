package logic

import (
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/0x587/guardeye/common/ent/agent"
	"github.com/0x587/guardeye/common/rediskey"
	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"

	"github.com/zeromicro/go-zero/core/logx"
)

type HeartbeatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHeartbeatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HeartbeatLogic {
	return &HeartbeatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HeartbeatLogic) Heartbeat(in *report.HeartbeatReq) (*report.Empty, error) {
	nodeInfo := in.GetNodeInfo()
	cid, err := uuid.Parse(nodeInfo.GetClientId())
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.Db.Agent.Update().
		SetIps(nodeInfo.GetNodeDescription().GetIps()).
		SetMacs(nodeInfo.GetNodeDescription().GetMacs()).
		SetOs(nodeInfo.GetNodeDescription().GetOs()).
		SetOsVersion(nodeInfo.GetNodeDescription().GetOsVersion()).
		SetHostname(nodeInfo.GetNodeDescription().GetHostname()).
		SetCPU(nodeInfo.GetNodeDescription().GetCpu()).
		SetMemory(nodeInfo.GetNodeDescription().GetMemory()).
		SetDisk(nodeInfo.GetNodeDescription().GetDisk()).
		SetUptime(nodeInfo.GetNodeDescription().GetUptime()).
		Where(agent.ClientIDEQ(cid)).
		Save(l.ctx)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.SetCtx(l.ctx, rediskey.SeeAtKey(nodeInfo), strconv.FormatInt(time.Now().UnixNano(), 10))
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.RedisClient.SetCtx(l.ctx, rediskey.LatencyKey(nodeInfo),
		strconv.FormatUint(uint64(in.GetLatency()), 10))
	if err != nil {
		return nil, err
	}
	return &report.Empty{}, nil
}
