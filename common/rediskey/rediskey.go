package rediskey

import (
	"fmt"

	"github.com/0x587/guardeye/report/report"
)

func TransDelayKey(nodeInfo *report.NodeInfo) string {
	return fmt.Sprintf("trans-delay-%s", nodeInfo.GetClientId())
}

func SeeAtKey(nodeInfo *report.NodeInfo) string {
	return fmt.Sprintf("see-at-%s", nodeInfo.GetClientId())
}

func NodeDescKey(nodeInfo *report.NodeInfo) string {
	return fmt.Sprintf("node-desc-%s", nodeInfo.GetClientId())
}

func LogDataKey(nodeInfo *report.NodeInfo) string {
	return fmt.Sprintf("log-data-key-%s", nodeInfo.GetClientId())
}

func LatencyKey(nodeInfo *report.NodeInfo) string {
	return fmt.Sprintf("latency-key-%s", nodeInfo.GetClientId())
}
