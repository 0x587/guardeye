package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"

	"github.com/0x587/guardeye/test-client/reporter/ros/pb/pb"
)

type impl struct {
}

func newImpl() (*impl, error) {

	return &impl{}, nil
}

type CallReq struct {
	Cid    string `json:"cid"`
	Action string `json:"action"`
	Method string `json:"method"`
	Data   string `json:"data"`
}
type CallRsp struct {
	Data string `json:"data"`
}

func (i *impl) Invoke(ctx context.Context, method string, args any, reply any, opts ...grpc.CallOption) error {
	data, err := json.Marshal(args)
	if err != nil {
		return err
	}
	req := &CallReq{
		Cid:    "106bc8ee-6048-4024-8afb-c294ec8fd559",
		Action: "send_topic",
		Method: method,
		Data:   string(data),
	}
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	postRsp, err := http.Post("http://localhost:8888/api/v1/downstream/call", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	rspBody, err := io.ReadAll(postRsp.Body)
	if err != nil {
		return err
	}
	defer postRsp.Body.Close()
	rsp := &CallRsp{}
	if err := json.Unmarshal(rspBody, rsp); err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(rsp.Data), reply); err != nil {
		return err
	}
	return nil
}

func (i *impl) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	//TODO implement me
	panic("implement me")
}

func main() {
	cc, err := newImpl()
	if err != nil {
		panic(err)
	}
	cli := pb.NewApiClient(cc)
	for {
		go func() {
			lo.Try(func() error {
				ctx := context.Background()
				res, err := cli.PublishTopicChatter(ctx, &pb.String{
					Data: time.Now().Format(time.RFC3339Nano),
				})
				if err != nil {
					panic(err)
				}
				t, _ := time.Parse(time.RFC3339Nano, res.Data)
				logx.Infof("%v", time.Now().Sub(t))
				return nil
			})
		}()
		time.Sleep(time.Second)
	}
}
