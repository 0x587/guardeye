package logic

import (
	"context"

	"github.com/0x587/guardeye/link/internal/svc"
	"github.com/0x587/guardeye/link/link"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAgentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AgentListLogic {
	return &AgentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AgentListLogic) AgentList(in *link.Empty) (*link.AgentListRsp, error) {
	rsp, err := agentConn.List(l.ctx)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}
