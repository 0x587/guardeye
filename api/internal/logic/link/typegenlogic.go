package link

import (
	"context"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/link/linkclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TypeGenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTypeGenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TypeGenLogic {
	return &TypeGenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TypeGenLogic) TypeGen(req *types.TypeGenReq) (resp *types.TypeGenRsp, err error) {
	rsp, err := l.svcCtx.LinkCli.TypeGen(l.ctx, &linkclient.TypeGenReq{
		Cid: req.Cid,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.TypeGenRsp{
		Pb: rsp.Pb,
	}
	return
}
