package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/0x587/guardeye/common/rediskey"
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
	logBytes, err := json.Marshal(in)
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
			ReceiveDelay: serverTime.UnixNano() - clientSendTime.UnixNano(),
		}
		bytes, err := json.Marshal(d)
		if err != nil {
			return nil, err
		}
		if err := l.svcCtx.RedisClient.SetCtx(l.ctx, rediskey.TransDelayKey(in.GetNodeInfo()), string(bytes)); err != nil {
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
