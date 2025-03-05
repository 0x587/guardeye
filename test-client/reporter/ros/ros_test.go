package ros

import (
	"os"
	"testing"
)

func Test_1(t *testing.T) {
	i := newImpl()
	pb := NewPbEnv()

	msg, err := i.getInterfaceMsgDefine("tutorial_interfaces/msg/Msg1")
	if err != nil {
		t.Fatal(err)
	}
	pb.AddMsg(&Message{
		Topic: "/msg",
		Msg:   msg,
	})

	msg, err = i.getInterfaceMsgDefine("std_msgs/msg/String")
	if err != nil {
		t.Fatal(err)
	}
	pb.AddMsg(&Message{
		Topic: "/chatter",
		Msg:   msg,
	})

	srv, err := i.getInterfaceSrvDefine("shawn_define/srv/Add")
	if err != nil {
		t.Fatal(err)
	}
	pb.AddSrv(&Service{
		Topic: "/add_two_ints_srv",
		Srv:   srv,
	})

	r := pb.Output()
	for k, v := range r {
		open, err := os.Create("pb/" + k + ".proto")
		if err != nil {
			t.Fatal(err)
		}
		_, err = open.WriteString(v)
		if err != nil {
			t.Fatal(err)
		}
		err = open.Close()
		if err != nil {
			t.Fatal(err)
		}
	}
}
