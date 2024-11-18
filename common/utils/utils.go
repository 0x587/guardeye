package utils

import (
	"fmt"
	"strings"

	"github.com/0x587/guardeye/report/report"
)

func ProviderToStr(p *report.Provider) string {
	return fmt.Sprintf("%s(%s)", p.GetType(), strings.Join(p.GetArgs(), ","))
}
