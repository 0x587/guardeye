package rostopic

import (
	"bufio"
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/0x587/guardeye/report/report"
	"github.com/0x587/guardeye/report/reportclient"
	"github.com/0x587/guardeye/test-client/provider"
	"github.com/zeromicro/go-zero/core/logc"
)

func New(ctx context.Context, topic string) provider.IF {
	res := &impl{
		topic: topic,
		out:   make(chan string),
	}
	go res.loop(ctx)
	return res
}

type impl struct {
	topic string
	out   chan string
}

func (i *impl) Get() <-chan *provider.Msg {
	res := make(chan *provider.Msg)
	go func() {
		for msg := range i.out {
			res <- &provider.Msg{
				Message:  msg,
				Type:     report.LogType_YAML,
				Provider: i.getProvider(),
			}
		}
	}()
	return res
}

func (i *impl) loop(ctx context.Context) {
	cmd := exec.Command("ros2", "topic", "echo", i.topic, "--full-length")
	stdout, err := cmd.StdoutPipe()
	logc.Must(err)
	go func() { logc.Must(cmd.Run()) }()
	scanner := bufio.NewScanner(stdout)
	topicCh := make(chan []string)
	defer close(topicCh)
	topicBlock := make([]string, 0)
	go func() {
		for topic := range topicCh {
			i.out <- strings.Join(topic, "\n")
		}
	}()
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
			if scanner.Text() == "---" {
				topicCh <- topicBlock
				topicBlock = make([]string, 0)
				continue
			}
			topicBlock = append(topicBlock, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from scanner:", err)
	}
}

const ProviderType = "RosTopic"

func (i *impl) getProvider() reportclient.Provider {
	return reportclient.Provider{
		Type: ProviderType,
		Args: []string{i.topic},
	}
}
