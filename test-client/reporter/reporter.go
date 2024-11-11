package reporter

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"

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

func New(cli zrpc.Client, providers ...provider.IF) IF {
	res := &impl{
		providers:    providers,
		reportClient: reportclient.NewReport(cli),
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
	cs := lo.Map(i.providers, func(item provider.IF, _ int) <-chan *provider.Msg {
		return item.Get()
	})
	for msg := range lo.FanIn(0, cs...) {
		logx.Infof("report: %v", msg)
		err := i.doLogReport(ctx, msg)
		if err != nil {
			logx.Errorf("report error: %v", err)
		}
	}
}

type clientConfig struct {
	ClientID string `json:"client_id"`
}

func (i *impl) doInit(ctx context.Context) error {
	f := lo.Must(os.OpenFile("./client.conf", os.O_CREATE|os.O_RDWR, 0666))
	defer func(f *os.File) {
		err := f.Close()
		logx.Must(err)
	}(f)
	c := &clientConfig{}
	err := json.NewDecoder(f).Decode(c)
	fmt.Printf("client config: %v\n", c)
	if err != nil || c.ClientID == "" {
		resp, err := i.reportClient.Init(ctx, &reportclient.InitReq{
			NodeDescription: i.getNodeDesc(),
		})
		if err != nil {
			return err
		}
		c.ClientID = resp.GetNodeInfo().GetClientId()
	}
	i.clientID = c.ClientID
	lo.Must(f.Seek(0, 0))
	lo.Must0(json.NewEncoder(f).Encode(c))
	return nil
}

func (i *impl) doLogReport(ctx context.Context, msg *provider.Msg) error {
	resp, err := i.reportClient.LogReport(ctx, &reportclient.LogReportReq{
		NodeInfo: &reportclient.NodeInfo{
			ClientId:        i.clientID,
			NodeDescription: i.getNodeDesc(),
		},
		Message:  msg.Message,
		Provider: &msg.Provider,
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
