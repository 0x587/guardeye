package link

import (
	"context"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"
	"github.com/0x587/guardeye/link/linkclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TypeListLogic {
	return &TypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TypeListLogic) TypeList(req *types.TypeListReq) (resp *types.TypeListRsp, err error) {
	rsp, err := l.svcCtx.LinkCli.TypeList(l.ctx, &linkclient.TypeListReq{
		Cid: req.Cid,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.TypeListRsp{
		Messages: rsp.Messages,
		Services: rsp.Services,
	}
	return
}
