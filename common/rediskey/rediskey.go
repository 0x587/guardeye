package rediskey

import (
	"fmt"

	"github.com/0x587/guardeye/report/report"
)

func TransDelayKey(nodeInfo *report.NodeInfo) string {
	return fmt.Sprintf("trans-delay-%s", nodeInfo.GetClientId())
}
