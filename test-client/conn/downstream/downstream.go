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
	var rosImpl ros.IF
	if c.Bridge.Enable {
		rosImpl = implfg.New(c.Bridge.Ip, c.Bridge.Port)
	}
	res := &mqttImpl{
		ctx:        ctx,
		cancelFunc: cancelFunc,
		cid:        cid.String(),
		rosImpl:    rosImpl,
	}
	if err := sLink.Init(res.callback); err != nil {
		return nil, err
	}
	return res, nil
}

type mqttImpl struct {
	ctx        context.Context
	cancelFunc context.CancelFunc

	sLink   serverLink
	cid     string
	rosImpl ros.IF
}

func (i *mqttImpl) callback(req *linkclient.LinkCommandDownstream) (*linkclient.LinkCommandUpstream, error) {
	rsp := &linkclient.LinkCommandUpstream{
		Id:     req.Id,
		Ok:     false,
		ErrMsg: "unknown type",
	}

	if req.GetPayloadRosExec() != nil {
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

func (i *mqttImpl) Close() {
	i.cancelFunc()
	i.sLink.Close()
}
