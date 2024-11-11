package provider

import "github.com/0x587/guardeye/report/reportclient"

type IF interface {
	Get() <-chan *Msg
}

type Msg struct {
	Message  string
	Provider reportclient.Provider
}
