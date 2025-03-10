package walk

import (
	"testing"

	"github.com/0x587/guardeye/test-client/conn/downstream/ros/implcli/rosmsg"
)

func TestToPb(t *testing.T) {
	msg := `
Vector3  linear
Vector3  angular`
	tree, err := rosmsg.ParseMsg(msg)
	if err != nil {
		t.Fatal(err)
	}
	Walk(tree)
	//type args struct {
	//	tree antlr.Tree
	//}
	//tests := []struct {
	//	name string
	//	args args
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		ToPb(tt.args.tree)
	//	})
	//}
}
