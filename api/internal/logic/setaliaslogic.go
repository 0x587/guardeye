package logic

import (
	"context"
	"database/sql"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
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
	node, err := l.svcCtx.NodeDBClient.FindOneLast(l.ctx, cid)
	if err != nil {
		return nil, err
	}
	node.Alias = sql.NullString{Valid: true, String: req.Alias}
	if err = l.svcCtx.NodeDBClient.Update(l.ctx, node); err != nil {
		return nil, err
	}
	return &types.SetAliasRsp{Ok: true}, nil
}
