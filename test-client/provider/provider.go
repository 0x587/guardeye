package provider

import (
	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/report/reportclient"
)

type IF interface {
	Get() <-chan *Msg
}

type Msg struct {
	Message  string
	Type     report.LogType
	Provider reportclient.Provider
}
