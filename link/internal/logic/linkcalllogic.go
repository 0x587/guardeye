package logic

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/0x587/guardeye/common/eventpool"
	"github.com/0x587/guardeye/common/foxglovetopb"
	"github.com/0x587/guardeye/foxglove_cdrservice/proto/foxgloveService"
	"github.com/0x587/guardeye/link/internal/svc"
	"github.com/0x587/guardeye/link/link"
	"github.com/0x587/guardeye/test-client/conn/downstream/ros"
	"github.com/0x587/guardeye/test-client/conn/downstream/ros/implcli"

	"github.com/zeromicro/go-zero/core/logx"
)

type LinkCallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var agentConn = newRpcAgentConn()
var listenPool = eventpool.New[string, *link.AgentListenRsp]()

func NewLinkCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LinkCallLogic {
	return &LinkCallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LinkCallLogic) LinkCall(req *link.LinkCallReq) (*link.LinkCallRsp, error) {
	action, topic, err := methodParse(req.GetMethod())
	if err != nil {
		return nil, err
	}

	schema, err := l.getSchema(req.GetCid(), action, topic)
	if err != nil {
		return nil, err
	}
	pbEnv := foxglovetopb.New()
	var pbReqTypeName, pbReqSchema string
	var pbRspTypeName, pbRspSchema string
	if action == ros.ActionCallService {
		if err := pbEnv.ParseSrv(topic, schema.GetName(), schema.GetReq(), schema.GetRsp()); err != nil {
			return nil, err
		}
		pbReqTypeName, pbReqSchema = pbEnv.OutputForType(schema.GetName() + "Req")
		pbRspTypeName, pbRspSchema = pbEnv.OutputForType(schema.GetName() + "Rsp")
	}
	if action == ros.ActionSendTopic {
		if err := pbEnv.ParseMsg(topic, schema.GetName(), schema.GetReq()); err != nil {
			return nil, err
		}
		pbReqTypeName, pbReqSchema = pbEnv.OutputForType(schema.GetName())
		pbRspTypeName, pbRspSchema = pbEnv.OutputForType(schema.GetName())
	}

	pbDataBytes, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		return nil, err
	}

	cdrWriteRsp, err := l.svcCtx.CdrCli.CdrWrite(l.ctx, &foxgloveService.CdrWriteReq{
		RosSchema:  schema.GetReq(),
		PbSchema:   pbReqSchema,
		PbTypeName: pbReqTypeName,
		TransData:  pbDataBytes,
	})
	if err != nil {
		return nil, err
	}

	cdrDataBytes := cdrWriteRsp.GetCdrData()
	cdrDataBase64 := base64.StdEncoding.EncodeToString(cdrDataBytes)

	p := &link.LinkCommandDownstream{
		PayloadRosExec: &link.LinkCommandPayloadRosExec{
			Action:   action,
			RosTopic: topic,
			Data:     cdrDataBase64,
		},
	}

	agentDoRsp, err := agentConn.Call(l.ctx, req.Cid, p)
	if err != nil {
		return nil, err
	}

	if !agentDoRsp.Ok {
		return nil, errors.New(agentDoRsp.ErrMsg)
	}
	cdrDataBytes, err = base64.StdEncoding.DecodeString(agentDoRsp.CdrData)
	if err != nil {
		return nil, err
	}
	cdrReadRsp, err := l.svcCtx.CdrCli.CdrRead(l.ctx, &foxgloveService.CdrReadReq{
		RosSchema:  schema.GetRsp(),
		PbSchema:   pbRspSchema,
		PbTypeName: pbRspTypeName,
		CdrData:    cdrDataBytes,
	})
	if err != nil {
		return nil, err
	}
	if action == ros.ActionSendTopic {
		listenPool.Invoke(req.GetCid(), &link.AgentListenRsp{
			Id:        uuid.New().String(),
			Timestamp: time.Now().UnixMilli(),
			Type:      link.AgentListenType_Download,
			Topic:     topic,
			Req:       cdrWriteRsp.GetJsonData(),
		})
	}
	if action == ros.ActionCallService {
		listenPool.Invoke(req.GetCid(), &link.AgentListenRsp{
			Id:        uuid.New().String(),
			Timestamp: time.Now().UnixMilli(),
			Type:      link.AgentListenType_Request,
			Topic:     topic,
			Req:       cdrWriteRsp.GetJsonData(),
			Rsp:       cdrReadRsp.GetJsonData(),
		})
	}

	readRspDataStr := base64.StdEncoding.EncodeToString(cdrReadRsp.GetTransData())

	return &link.LinkCallRsp{
		Data: readRspDataStr,
	}, nil
}

func (l *LinkCallLogic) getSchema(cid, action, topic string) (*link.LinkTypeGenResult, error) {
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

func methodParse(method string) (action string, topic string, err error) {
	method = strings.Replace(method, "/Api/", "", 1)
	if strings.HasPrefix(method, "PublishTopic") {
		topic = strings.Replace(method, "PublishTopic", "", 1)
		topic = implcli.Name2Topic(topic)
		action = ros.ActionSendTopic
	} else if strings.HasPrefix(method, "CallService") {
		topic = strings.Replace(method, "CallService", "", 1)
		topic = implcli.Name2Topic(topic)
		action = ros.ActionCallService
	} else {
		return "", "", errors.New(fmt.Sprintf("unknown method %s", method))
	}
	return action, topic, nil
}
