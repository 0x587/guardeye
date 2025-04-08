package logic

import (
	"context"

	"github.com/0x587/guardeye/link/internal/svc"
	"github.com/0x587/guardeye/link/link"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgentListenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAgentListenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AgentListenLogic {
	return &AgentListenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AgentListenLogic) AgentListen(in *link.AgentListenReq, stream link.Link_AgentListenServer) error {
	cid := in.GetCid()
	for {
		listenRsp, err := listenPool.Wait(l.ctx, cid)
		if err != nil {
			return err
		}
		if err := stream.Send(listenRsp); err != nil {
			return err
		}
	}
}
