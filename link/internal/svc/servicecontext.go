package svc

import (
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/0x587/guardeye/foxglove_cdrservice/proto/foxgloveService"
	"github.com/0x587/guardeye/link/internal/config"
)

type ServiceContext struct {
	Config config.Config
	CdrCli foxgloveService.FoxgloveServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	cdrClient := lo.Must(grpc.NewClient("127.0.0.1:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials())))
	cdrCli := foxgloveService.NewFoxgloveServiceClient(cdrClient)
	return &ServiceContext{
		Config: c,
		CdrCli: cdrCli,
	}
}
