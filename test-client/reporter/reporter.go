package reporter

import (
	"context"
	"net"
	"os"
	"runtime"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/report/reportclient"
	"github.com/0x587/guardeye/test-client/feature"
	"github.com/0x587/guardeye/test-client/feature/featuredelay"
	"github.com/0x587/guardeye/test-client/provider"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type IF interface {
	Loop(ctx context.Context)
}

func New(providers ...provider.IF) IF {
	clientConf := zrpc.RpcClientConf{
		Target: "localhost:8080",
	}
	conn := zrpc.MustNewClient(clientConf)
	res := &impl{
		providers:    providers,
		reportClient: reportclient.NewReport(conn),
		featureDelay: featuredelay.New(),
	}
	logx.Must(res.doInit(context.Background()))
	return res
}

type impl struct {
	clientID     string
	desc         *reportclient.NodeDescription
	reportClient reportclient.Report
	featureDelay feature.IF[
		*reportclient.FeatureTransDelayReq,
		*reportclient.FeatureTransDelayRsp,
	]
	providers []provider.IF
}

func (i *impl) Loop(ctx context.Context) {
	cs := lo.Map(i.providers, func(item provider.IF, _ int) <-chan string {
		return item.Get()
	})
	for log := range lo.FanIn(0, cs...) {
		err := i.doLogReport(ctx, log)
		if err != nil {
			logx.Errorf("report error: %v", err)
		}
	}
}

func (i *impl) doInit(ctx context.Context) error {
	resp, err := i.reportClient.Init(ctx, &reportclient.InitReq{
		NodeDescription: i.getNodeDesc(),
	})
	if err != nil {
		return err
	}
	i.clientID = resp.GetNodeInfo().GetClientId()
	return nil
}

func (i *impl) doLogReport(ctx context.Context, msg string) error {
	resp, err := i.reportClient.LogReport(ctx, &reportclient.LogReportReq{
		NodeInfo: &reportclient.NodeInfo{
			ClientId:        i.clientID,
			NodeDescription: i.getNodeDesc(),
		},
		Level:   report.LogLevel_DEBUG,
		Message: msg,
		Features: &reportclient.FeaturesReq{
			TransDelay: lo.Must(i.featureDelay.MakeReq()),
		},
	})
	if err != nil {
		return err
	}
	lo.Must0(i.featureDelay.HandleRsp(resp.GetFeatures().GetTransDelay()))
	return nil
}

func (i *impl) getNodeDesc() *reportclient.NodeDescription {
	addr := lo.Must(net.InterfaceAddrs())
	interfaces := lo.Must(net.Interfaces())
	return &reportclient.NodeDescription{
		Os:        runtime.GOOS,
		OsVersion: runtime.GOARCH,
		Alias:     "test-client",
		Ips: lo.Map(addr, func(item net.Addr, _ int) string {
			return item.String()
		}),
		Macs: lo.Map(interfaces, func(item net.Interface, _ int) string {
			return item.HardwareAddr.String()
		}),
		Hostname: lo.Must(os.Hostname()),
	}
}
