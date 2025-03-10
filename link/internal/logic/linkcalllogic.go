package logic

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logc"

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
	fmt.Printf("%s\n", schema)

	pbDataBytes, err := base64.StdEncoding.DecodeString(req.Data)
	if err != nil {
		return nil, err
	}
	cdrWriteRsp, err := l.svcCtx.CdrCli.CdrWrite(l.ctx, &foxgloveService.CdrWriteReq{
		RosSchema: `
V2 a
V2 b
int64[] c
float64 d
float64[] e
================================================================================
MSG: shawn_define/V2
int64 a
int64 b`,
		PbSchema: `
syntax = "proto3";
package shawn_define;

message AddReq {
	V2 a = 1;
	V2 b = 2;
	repeated int64 c = 3;
	double d = 4;
	repeated double e = 5;
}

message V2 {
	int64 a = 1;
	int64 b = 2;
}
`, // TODO
		PbTypeName: "shawn_define.AddReq", // TODO
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

	logc.Infof(l.ctx, "send to %s", req.Cid)
	agentDoRsp, err := agentConn.Call(l.ctx, req.Cid, p)
	if err != nil {
		return nil, err
	}

	if !agentDoRsp.Ok {
		return nil, errors.New(agentDoRsp.ErrMsg)
	}
	logc.Infof(l.ctx, "revice from %s", req.Cid)
	cdrDataBytes, err = base64.StdEncoding.DecodeString(agentDoRsp.CdrData)
	if err != nil {
		return nil, err
	}
	cdrReadRsp, err := l.svcCtx.CdrCli.CdrRead(l.ctx, &foxgloveService.CdrReadReq{
		RosSchema: `
V2 res
int64[] c
float64 d
float64[] e
================================================================================
MSG: shawn_define/V2
int64 a
int64 b
`, // TODO
		PbSchema: `
syntax = "proto3";
package shawn_define;

message V2 {
	int64 a = 1;
	int64 b = 2;
}

message AddRsp {
	V2 res = 1;
	repeated int64 c = 2;
	double d = 3;
	repeated double e = 4;
}
`, // TODO
		PbTypeName: "shawn_define.AddRsp", // TODO
		CdrData:    cdrDataBytes,
	})
	if err != nil {
		return nil, err
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
