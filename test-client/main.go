package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/0x587/guardeye/test-client/provider/filewatch"
	"github.com/0x587/guardeye/test-client/reporter"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	fileProvider := filewatch.New(ctx, "./test")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	r := reporter.New(fileProvider)
	go r.Loop(ctx)
	<-sigChan
	cancel()
}
