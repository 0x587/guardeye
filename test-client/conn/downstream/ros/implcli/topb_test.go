package implcli

import (
	"fmt"
	"os"
	"testing"

	"github.com/samber/lo"

	"github.com/0x587/guardeye/test-client/conn/downstream/ros"
)

func TestNewPbEnv(t *testing.T) {
	env := NewPbEnv()
	i := &impl{}
	env.AddSrv(&ros.Service{
		Topic: "/add_two_ints_srv",
		Srv:   lo.Must(i.getInterfaceSrvDefine("shawn_define/srv/Add")),
	})

	a := env.Output()
	for name, content := range a {
		f := lo.Must(os.Create(fmt.Sprintf("./pb/%s.proto", name)))
		f.WriteString(content)
	}
}
