package logic

import (
	"context"

	"google.golang.org/protobuf/proto"

	"github.com/0x587/guardeye/hub/internal/svc"
	"github.com/0x587/guardeye/hub/report"
	"github.com/0x587/guardeye/report/reportclient"

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
	req := proto.Clone(in).(*reportclient.InitReq)
	out, err := l.svcCtx.Upstream.Init(l.ctx, req)
	rsp := proto.Clone(out).(*report.InitRsp)
	return rsp, err
}
