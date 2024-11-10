package main

import (
	"context"
	"log"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/report/reportclient"
	"github.com/0x587/guardeye/test-client/feature"
	"github.com/0x587/guardeye/test-client/feature/featuredelay"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/zrpc"
)

func main() {
	i := NewImpl()
	err := i.doInit(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for {
		err := i.doLogReport(
			context.Background(),
		)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}

type impl struct {
	clientID     string
	desc         *reportclient.NodeDescription
	reportClient reportclient.Report
	featureDelay feature.IF[
		*reportclient.FeatureTransDelayReq,
		*reportclient.FeatureTransDelayRsp,
	]
}

func NewImpl() *impl {
	clientConf := zrpc.RpcClientConf{
		//Etcd: discov.EtcdConf{},
		//Endpoints:     nil,
		Target: "localhost:8080",
		//App:           "",
		//Token:         "",
		//NonBlock:      false,
		//Timeout:       0,
		//KeepaliveTime: 0,
		//Middlewares:   zrpc.ClientMiddlewaresConf{},
	}
	conn := zrpc.MustNewClient(clientConf)
	return &impl{
		reportClient: reportclient.NewReport(conn),
		featureDelay: featuredelay.New(),
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

func (i *impl) doLogReport(ctx context.Context) error {
	resp, err := i.reportClient.LogReport(ctx, &reportclient.LogReportReq{
		NodeInfo: &reportclient.NodeInfo{
			ClientId:        i.clientID,
			NodeDescription: i.getNodeDesc(),
		},
		Level:   report.LogLevel_DEBUG,
		Message: time.Now().String(),
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
		OsVersion: runtime.Version(),
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
