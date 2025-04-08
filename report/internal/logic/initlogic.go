package logic

import (
	"context"
	"errors"

	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"

	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitLogic {
	return &InitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitLogic) Init(in *report.InitReq) (*report.InitRsp, error) {
	if in.GetNodeDescription() == nil {
		return nil, errors.New("missing node description")
	}
	uid := uuid.New()
	a, err := l.svcCtx.Db.Agent.Create().
		SetClientID(uid).
		SetIps(in.GetNodeDescription().GetIps()).
		SetMacs(in.GetNodeDescription().GetMacs()).
		SetOs(in.GetNodeDescription().GetOs()).
		SetOsVersion(in.GetNodeDescription().GetOsVersion()).
		SetHostname(in.GetNodeDescription().GetHostname()).
		SetCPU(in.GetNodeDescription().GetCpu()).
		SetMemory(in.GetNodeDescription().GetMemory()).
		SetDisk(in.GetNodeDescription().GetDisk()).
		SetUptime(in.GetNodeDescription().GetUptime()).
		Save(l.ctx)
	if err != nil {
		return nil, err
	}
	logx.Info(a)
	return &report.InitRsp{
		NodeInfo: &report.NodeInfo{
			ClientId:        uid.String(),
			NodeDescription: in.GetNodeDescription(),
		},
	}, nil
}
