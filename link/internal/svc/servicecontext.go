package svc

import (
	_ "github.com/lib/pq"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/0x587/guardeye/common/ent"
	"github.com/0x587/guardeye/common/httpcb"
	"github.com/0x587/guardeye/foxglove_cdrservice/proto/foxgloveService"
	"github.com/0x587/guardeye/link/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	CdrCli       foxgloveService.FoxgloveServiceClient
	HttpCallback httpcb.IF
}

func NewServiceContext(c config.Config) *ServiceContext {
	cdrClient := lo.Must(grpc.NewClient(c.CdrServiceTarget,
		grpc.WithTransportCredentials(insecure.NewCredentials())))
	cdrCli := foxgloveService.NewFoxgloveServiceClient(cdrClient)
	db := lo.Must(ent.Open(c.PostgresConf.Driver, c.PostgresConf.Dsn))
	return &ServiceContext{
		Config:       c,
		CdrCli:       cdrCli,
		HttpCallback: httpcb.New(db),
	}
}
