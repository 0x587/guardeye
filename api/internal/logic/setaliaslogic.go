package logic

import (
	"context"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
