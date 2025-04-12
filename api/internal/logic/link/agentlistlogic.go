package link

import (
	"context"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/link/linkclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AgentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAgentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AgentListLogic {
	return &AgentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AgentListLogic) AgentList(req *types.AgentListReq) (resp *types.AgentListRsp, err error) {
	rsp, err := l.svcCtx.LinkCli.AgentList(l.ctx, &linkclient.Empty{})
	if err != nil {
		return nil, err
	}
	resp = &types.AgentListRsp{
		Agents: rsp.GetAgents(),
	}
	return
}
