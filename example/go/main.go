package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/0x587/guardeye/example/go/pb/pb"
	"github.com/0x587/guardeye/sdk/gosdk"
)

func main() {
	client := gosdk.NewClientConn("localhost:8080")
	cli := pb.NewApiClient(client)
	for {
		ctx := context.Background()
		person, err := cli.CallServiceGetPerson(ctx, &pb.GetPersonReq{
			Req: &pb.Person{
				V: &pb.V3{X: 1, Y: 2, Z: 3},
				Pos: &pb.Pos{
					Type: "123",
					V:    &pb.V3{X: 1, Y: 2, Z: 3},
				},
				State: &pb.State{
					Count: 213,
					V:     &pb.V3{X: 1, Y: 2, Z: 3},
				},
			},
		})
		if err != nil {
			return
		}
		personJson, _ := json.Marshal(person)
		fmt.Printf("%s\n", personJson)
		time.Sleep(time.Second)
	}

}
