package logic

import (
	"context"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/common/ent/agent"

	"github.com/google/uuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetAliasLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetAliasLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAliasLogic {
	return &SetAliasLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetAliasLogic) SetAlias(req *types.SetAliasReq) (resp *types.SetAliasRsp, err error) {
	cid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.Db.Agent.Update().Where(agent.ClientIDEQ(cid)).SetAlias(req.Alias).Save(l.ctx)
	if err != nil {
		return nil, err
	}
	return &types.SetAliasRsp{Ok: true}, nil
}
