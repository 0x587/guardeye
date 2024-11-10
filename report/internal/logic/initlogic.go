package logic

import (
	"context"

	model "github.com/0x587/guardeye/report/internal/models"
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
	uid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.NodeClient.Insert(l.ctx, &model.Node{
		Description: *in.GetNodeDescription(),
		Ip:          in.GetNodeDescription().GetIp(),
	})
	if err != nil {
		return nil, err
	}
	return &report.InitRsp{
		NodeInfo: &report.NodeInfo{
			ClientId:        uid.String(),
			NodeDescription: in.GetNodeDescription(),
		},
	}, nil
}
