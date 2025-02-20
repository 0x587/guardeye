package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/0x587/guardeye/common/rediskey"
	"github.com/0x587/guardeye/common/tokv"
	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"

	"github.com/golang/protobuf/jsonpb"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/yaml.v3"
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
		parsed, parsedType := parse(log.GetMessage())
		mqLog := &report.MQLog{
			Timestamp: time.UnixMilli(int64(log.ReportAtMilli)).Format(time.RFC3339),
			NodeInfo:  in.GetNodeInfo(),
			Log: &report.Log{
				Message:  log.GetMessage(),
				Provider: log.GetProvider(),
				Type:     log.GetType(),
			},
			Parsed: parsed,
			Flag: &report.MQLogFlag{
				JsonAble: parsedType == JsonAble,
				YamlAble: parsedType == YamlAble,
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

	// set see at
	if err := l.svcCtx.RedisClient.SetCtx(l.ctx, rediskey.SeeAtKey(in.GetNodeInfo()),
		fmt.Sprintf("%d", time.Now().UnixNano())); err != nil {
		return nil, err
	}

	//update node description
	//desc := in.GetNodeInfo().GetNodeDescription()

	return rsp, nil
}

const (
	JsonAble = "json"
	YamlAble = "yaml"
)

func parse(v string) (map[string]string, string) {
	var obj interface{}
	if err := json.Unmarshal([]byte(v), &obj); err == nil {
		return tokv.ObjToKv(obj), JsonAble
	}
	obj = struct{}{}
	if err := yaml.Unmarshal([]byte(v), &obj); err == nil {
		return tokv.ObjToKv(obj), YamlAble
	}
	return make(map[string]string), ""
}
