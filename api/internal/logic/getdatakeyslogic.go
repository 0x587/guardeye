package logic

import (
	"context"

	"github.com/0x587/guardeye/api/internal/svc"
	"github.com/0x587/guardeye/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDataKeysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDataKeysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDataKeysLogic {
	return &GetDataKeysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDataKeysLogic) GetDataKeys(req *types.GetDataKeysReq) (resp *types.GetDataKeysRsp, err error) {
	// todo: add your logic here and delete this line

	return
}
