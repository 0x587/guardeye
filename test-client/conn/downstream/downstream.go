package downstream

import (
	"encoding/base64"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/net/context"

	"github.com/0x587/guardeye/link/linkclient"
	"github.com/0x587/guardeye/test-client/config"
	"github.com/0x587/guardeye/test-client/conn/downstream/ros"
	"github.com/0x587/guardeye/test-client/conn/downstream/ros/implfg"
)

type IF interface {
	Close()
}

func New(cid uuid.UUID, c config.Config) (IF, error) {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	sLink, err := newRpcServerLink(cid, c.LinkEndpoint)
	if err != nil {
		return nil, err
	}

	res := &impl{
		ctx:        ctx,
		cancelFunc: cancelFunc,
		cid:        cid.String(),
	}
	if c.Bridge.Enable {
		res.rosImpl = implfg.New(c.Bridge.Ip, c.Bridge.Port, c.Bridge.Patterns, res.subscribeTopicCallback)
	}
	if err := sLink.Init(res.callback); err != nil {
		return nil, err
	}
	res.sLink = sLink
	return res, nil
}

type impl struct {
	ctx        context.Context
	cancelFunc context.CancelFunc

	sLink   serverLink
	cid     string
	rosImpl ros.IF
}

func (i *impl) callback(req *linkclient.LinkCommandDownstream) (*linkclient.LinkCommandUpstream, error) {
	rsp := &linkclient.LinkCommandUpstream{
		Id:     req.Id,
		Ok:     false,
		ErrMsg: "unknown type",
	}

	if req.GetPayloadRosExec() != nil {
		// 当为订阅请求时 data字段填充轮训id
		if req.GetPayloadRosExec().GetAction() == ros.ActionSubscribeTopic {
			req.GetPayloadRosExec().Data = req.GetId()
		}
		res, err := i.rosImpl.Exec(req.GetPayloadRosExec())

		if err != nil {
			logx.Error(err)
			rsp.Ok = false
			rsp.ErrMsg = err.Error()
		} else {
			rsp.Ok = true
			rsp.CdrData = base64.StdEncoding.EncodeToString(res)
		}
		logx.Info("reply ", rsp)
	}
	if req.GetPayloadRosList() != nil {
		res, err := i.rosImpl.List(req.GetPayloadRosList())
		if err != nil {
			logx.Error(err)
			rsp.Ok = false
			rsp.ErrMsg = err.Error()
		} else {
			rsp.Ok = true
			rsp.TypeListResult = res
		}
	}
	if req.GetPayloadRosType() != nil {
		typeResult, err := i.rosImpl.Type(req.GetPayloadRosType())
		if err != nil {
			logx.Error(err)
			rsp.Ok = false
			rsp.ErrMsg = err.Error()
		} else {
			rsp.Ok = true
			rsp.TypeGenResult = typeResult
		}
	}

	return rsp, nil
}

func (i *impl) subscribeTopicCallback(id string, transData []byte) {
	_ = i.sLink.Send(&linkclient.LinkCommandUpstream{
		Id: id,
		TopicSubscribeResult: &linkclient.TopicSubscribeRsp{
			CdrData: base64.StdEncoding.EncodeToString(transData),
		},
	})
}

func (i *impl) Close() {
	i.cancelFunc()
	i.sLink.Close()
}
