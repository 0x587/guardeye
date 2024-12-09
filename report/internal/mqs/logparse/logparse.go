package logparse

import (
	"context"

	"github.com/zeromicro/go-queue/kq"

	"github.com/0x587/guardeye/report/internal/svc"
)

func New(ctx context.Context, svcCtx *svc.ServiceContext) kq.ConsumeHandler {
	return &impl{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type impl struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func (i impl) Consume(ctx context.Context, key, value string) error {
	//TODO implement me
	panic("implement me")
}
