package featuredelay

import (
	"time"

	"github.com/0x587/guardeye/report/reportclient"
	"github.com/0x587/guardeye/test-client/feature"
)

func New() feature.IF[
	*reportclient.FeatureTransDelayReq,
	*reportclient.FeatureTransDelayRsp,
] {
	return &impl{}
}

type impl struct {
	lastSend    int64
	lastReceive int64
}

func (i *impl) MakeReq() (*reportclient.FeatureTransDelayReq, error) {
	return &reportclient.FeatureTransDelayReq{
		Enable:               true,
		LastSendTimestamp:    i.lastSend,
		LastReceiveTimestamp: i.lastReceive,
		SentAtTimestamp:      time.Now().UnixNano(),
	}, nil
}

func (i *impl) HandleRsp(rsp *reportclient.FeatureTransDelayRsp) error {
	i.lastSend = rsp.GetSentAtTimestamp()
	i.lastReceive = time.Now().UnixNano()
	return nil
}
