package logic

import (
	"context"
	"io"
	"strconv"

	"github.com/0x587/guardeye/report/internal/svc"
	"github.com/0x587/guardeye/report/report"

	"github.com/zeromicro/go-zero/core/logx"
)

type LinkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLinkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LinkLogic {
	return &LinkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LinkLogic) Link(stream report.Report_LinkServer) error {
	// todo: add your logic here and delete this line
	n := 1
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = stream.Send(&report.LinkRsp{
			Answer: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Question,
		})
		if err != nil {
			return err
		}
		n++
		logx.Infof("from stream client question: %s", req.Question)
	}
}
