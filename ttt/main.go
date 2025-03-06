package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/0x587/guardeye/ttt/foxgloveclient"
)

type (
	M  map[string]any
	MS []M
)

func main() {
	logx.MustSetup(logx.LogConf{Encoding: "plain"})
	ctx := context.Background()
	fgCli := foxgloveclient.New("127.0.0.1", 8765)
	go func() {
		err := fgCli.Run(ctx)
		if err != nil {
			logx.Error(err)
		}
	}()
	time.Sleep(time.Second)
	//err := fgCli.Subscribe("/chatter", func(jsonData string) {
	//	logx.Infof(jsonData)
	//})
	//if err != nil {
	//	logx.Error(err)
	//}
	//go func() {
	//	for {
	//		time.Sleep(time.Second)
	//		err := fgCli.Publish("/chatter",
	//			fmt.Sprintf("{\"data\":\"%s\"}", time.Now().Format(time.RFC3339Nano)))
	//		if err != nil {
	//			logx.Error(err)
	//		}
	//	}
	//}()
	go func() {
		for {
			time.Sleep(time.Second)
			data := M{
				"a": M{
					"a": int64(100000000000000000),
					"b": 2,
				},
				"b": M{
					"a": 3,
					"b": 4,
				},
				"c": []int64{int64(100000000000000000), int64(200000000000000000)},
				"d": 5.87,
				"e": []float64{5.87, 6.87},
			}
			bytes, err := json.Marshal(data)
			if err != nil {
				return
			}
			rsp, err := fgCli.Call("/add_two_ints_srv", string(bytes))
			if err != nil {
				logx.Error(err)
			}
			//logx.Info(base64.StdEncoding.EncodeToString(rsp))
			logx.Info(rsp)
			var got M
			json.Unmarshal([]byte(rsp), &got)
			logx.Info(got)
			return
		}
	}()
	time.Sleep(time.Hour * 1000)
	//for m := range fgCli.GetOutputChan() {
	//	fmt.Println(m)
	//}
}
