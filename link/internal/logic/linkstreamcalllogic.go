package logic

import (
	"context"
	"encoding/base64"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/0x587/guardeye/common/foxglovetopb"
	"github.com/0x587/guardeye/foxglove_cdrservice/proto/foxgloveService"
	"github.com/0x587/guardeye/link/internal/svc"
	"github.com/0x587/guardeye/link/link"
	"github.com/0x587/guardeye/test-client/conn/downstream/ros"

	"github.com/zeromicro/go-zero/core/logx"
)

type LinkStreamCallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLinkStreamCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LinkStreamCallLogic {
	return &LinkStreamCallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LinkStreamCallLogic) LinkStreamCall(req *link.LinkCallReq, stream link.Link_LinkStreamCallServer) error {
	action, topic, err := methodParse(req.GetMethod())
	if err != nil {
		return err
	}
	schema, err := l.getSchema(req.GetCid(), action, topic)
	if err != nil {
		return err
	}
	pbEnv := foxglovetopb.New()
	var pbRspTypeName, pbRspSchema string
	if action == ros.ActionSubscribeTopic {
		if err := pbEnv.ParseMsg(topic, schema.GetName(), schema.GetReq()); err != nil {
			return err
		}
		pbRspTypeName, pbRspSchema = pbEnv.OutputForType(schema.GetName())
	} else {
		return errors.New("unknown action")
	}
	cid, err := uuid.Parse(req.GetCid())
	if err != nil {
		return err
	}

	// 发起订阅 请求终端上报数据
	callRsp, err := agentConn.Call(l.ctx, req.GetCid(), &link.LinkCommandDownstream{
		PayloadRosExec: &link.LinkCommandPayloadRosExec{
			Action:   action,
			RosTopic: topic,
		},
	})
	if err != nil {
		return err
	}

	loop := func() error {
		waitRsp, err := linkRpcEvent.Wait(l.ctx, callRsp.GetId())
		if err != nil {
			return err
		}

		cdrDataBytes, err := base64.StdEncoding.DecodeString(waitRsp.GetTopicSubscribeResult().GetCdrData())
		if err != nil {
			return err
		}
		cdrReadRsp, err := l.svcCtx.CdrCli.CdrRead(l.ctx, &foxgloveService.CdrReadReq{
			RosSchema:  schema.GetReq(),
			PbSchema:   pbRspSchema,
			PbTypeName: pbRspTypeName,
			CdrData:    cdrDataBytes,
		})
		if err != nil {
			return err
		}
		logx.Info(pbRspSchema, pbRspTypeName)
		logx.Info(cdrDataBytes)
		logx.Info(cdrReadRsp.GetTransData(), cdrReadRsp.GetJsonData())
		if action == ros.ActionSubscribeTopic {
			listenPool.Invoke(req.GetCid(), &link.AgentListenRsp{
				Id:        uuid.New().String(),
				Timestamp: time.Now().UnixMilli(),
				Type:      link.AgentListenType_Upload,
				Topic:     topic,
				Rsp:       cdrReadRsp.GetJsonData(),
			})
		}

		readRspDataStr := base64.StdEncoding.EncodeToString(cdrReadRsp.GetTransData())
		res := &link.LinkCallRsp{
			Data: readRspDataStr,
		}
		if req.GetNeedJson() {
			res.Json = cdrReadRsp.GetJsonData()
		}

		go l.svcCtx.HttpCallback.Data(l.ctx, cid, topic, cdrReadRsp.GetJsonData())
		if err := stream.Send(res); err != nil {
			return err
		}
		return nil
	}
	for {
		if err := loop(); err != nil {
			logx.Error(err)
			return err
		}
	}
}

func (l *LinkStreamCallLogic) getSchema(cid, action, topic string) (*link.LinkTypeGenResult, error) {
	t := link.RosTypeGenType_MESSAGE
	if action == ros.ActionCallService {
		t = link.RosTypeGenType_SERVICE
	}
	p := &link.LinkCommandDownstream{
		PayloadRosType: &link.LinkCommandPayloadRosType{
			Type:     t,
			RosTopic: topic,
		},
	}
	rsp, err := agentConn.Call(l.ctx, cid, p)
	if err != nil {
		return nil, err
	}
	return rsp.GetTypeGenResult(), nil
}
