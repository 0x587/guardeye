package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"
	"google.golang.org/protobuf/proto"

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
	// product raw log
	logBytes, err := proto.Marshal(in)
	if err != nil {
		return nil, err
	}
	if err := l.svcCtx.RawLogPusherClient.Push(l.ctx, string(logBytes)); err != nil {
		return nil, err
	}
	rsp := &report.LogReportRsp{
		Code: report.ReportResultCode_SUCCESS,
	}

	if in.GetFeatures() == nil {
		return rsp, nil
	}
	feats := in.GetFeatures()
	// calc delay
	if feats.GetTransDelay().GetEnable() {
		lastSendTime := time.Unix(0, feats.GetTransDelay().GetLastSendTimestamp())
		lastReceiveTime := time.Unix(0, feats.GetTransDelay().GetLastReceiveTimestamp())
		clientSendTime := time.Unix(0, feats.GetTransDelay().GetSentAtTimestamp())
		serverTime := time.Now()
		d := &report.DelayResult{
			SendDelay:    lastReceiveTime.UnixNano() - lastSendTime.UnixNano(),
			ReceiveDelay: clientSendTime.UnixNano() - serverTime.UnixNano(),
		}
		bytes, err := proto.Marshal(d)
		if err != nil {
			return nil, err
		}
		if err := l.svcCtx.RedisClient.SetCtx(l.ctx, delayKey(in.GetNodeInfo()), string(bytes)); err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func delayKey(nodeInfo *report.NodeInfo) string {
	return fmt.Sprintf("client-id-%s", nodeInfo.GetClientId())
}
