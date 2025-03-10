package rossrv

import (
	"testing"

	"github.com/0x587/guardeye/test-client/conn/downstream/ros/implcli/rossrv/walk"
)

func TestParse(t *testing.T) {
	s := `shawn_define.srv.Add_Response(res=shawn_define.msg.V2(a=3, b=7), magic=[5, 8, 7], s='587', ss=['5', '8', '7'])`
	tree, err := Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	res := walk.Walk(tree)
	t.Log(res)
}
