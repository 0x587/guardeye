package logic

import (
	"context"

	"google.golang.org/protobuf/proto"

	"github.com/0x587/guardeye/hub/internal/svc"
	"github.com/0x587/guardeye/hub/report"
	"github.com/0x587/guardeye/report/reportclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogReportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogReportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogReportLogic {
	return &LogReportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogReportLogic) LogReport(in *report.LogReportReq) (*report.LogReportRsp, error) {
	req := proto.Clone(in).(*reportclient.LogReportReq)
	out, err := l.svcCtx.Upstream.LogReport(l.ctx, req)
	rsp := proto.Clone(out).(*report.LogReportRsp)
	return rsp, err
}
