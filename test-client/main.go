package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/kardianos/service"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/test-client/reporter"
)

type program struct {
	cancel context.CancelFunc
}

func (p *program) Start(s service.Service) error {
	logx.SetWriter(logger)

	ctx := context.Background()
	ctx, p.cancel = context.WithCancel(ctx)
	c := LoadConfig(ctx)
	r := reporter.New(fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port), c.Providers...)
	go r.Loop(ctx)
	return nil
}

func (p *program) Stop(s service.Service) error {
	p.cancel()
	return nil
}

func main() {
	svcFlag := flag.String("service", "", "Control the system service.")
	flag.Parse()

	options := make(service.KeyValue)
	options["Restart"] = "always"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	svcConfig := &service.Config{
		Name:        "GuardEYE",
		DisplayName: "GuardEYE",
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target"},
		Option: options,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 10)
	l, err := s.Logger(errs)
	logger = &LogWriter{Logger: l}
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if len(*svcFlag) != 0 {
		err := service.Control(s, *svcFlag)
		if err != nil {
			log.Printf("Valid actions: %q\n", service.ControlAction)
			log.Fatal(err)
		}
		return
	}
	err = s.Run()
	if err != nil {
		logx.Error(err)
	}
}
