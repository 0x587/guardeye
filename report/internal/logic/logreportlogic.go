package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"time"

	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"

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
	if in.GetNodeInfo() == nil {
		return nil, errors.New(fmt.Sprintf("missing node info"))
	}
	// product raw log
	for _, log := range in.GetLogs() {
		t := time.UnixMilli(int64(log.ReportAtMilli))
		mqLog := &report.MQLog{
			Timestamp: t.Format(time.RFC3339),
			NodeInfo:  in.GetNodeInfo(),
			Log: &report.Log{
				Message:  log.GetMessage(),
				Provider: log.GetProvider(),
				Type:     log.GetType(),
			},
			Trace: []*report.TraceSpan{
				{
					Name:   "report",
					Action: "client report",
					Time:   uint64(time.Now().UnixNano()),
				},
			},
		}
		m := jsonpb.Marshaler{}
		losStr, err := m.MarshalToString(mqLog)
		if err != nil {
			return nil, err
		}
		if err := l.svcCtx.RawLogPusherClient.Push(l.ctx, losStr); err != nil {
			return nil, err
		}
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
			ReceiveDelay: serverTime.UnixNano() - clientSendTime.UnixNano(),
		}
		if err := l.svcCtx.DelayRedisClient.SetDelay(l.ctx, in.GetNodeInfo(), d); err != nil {
			return nil, err
		}
		rsp.Features = &report.FeaturesRsp{
			TransDelay: &report.FeatureTransDelayRsp{
				Enable:          true,
				SentAtTimestamp: time.Now().UnixNano(),
			},
		}
	}
	return rsp, nil
}
