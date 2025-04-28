package helper

import (
	"context"
	"errors"
	"fmt"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/link/linkclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RpcProxyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRpcProxyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RpcProxyLogic {
	return &RpcProxyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RpcProxyLogic) RpcProxy(req *types.RpcProxyReq) (resp *types.RpcProxyRsp, err error) {
	if !((req.Action == "PublishTopic") || (req.Action == "CallService")) {
		return nil, errors.New(fmt.Sprintf(
			"unsupport action %s not in ['PublishTopic', 'CallService']", req.Action))
	}
	rsp, err := l.svcCtx.LinkCli.LinkCall(l.ctx, &linkclient.LinkCallReq{
		Cid:    req.Cid,
		Method: req.Action + topic2Name(req.Topic),
		Data:   req.Data,
	})
	if err != nil {
		return nil, err
	}
	return &types.RpcProxyRsp{
		Data: rsp.GetData(),
	}, nil
}
