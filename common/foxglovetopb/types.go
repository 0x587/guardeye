package foxglovetopb

import "github.com/0x587/guardeye/common/foxglovetopb/walk"

type MsgDefine struct {
	Package string
	Name    string
	Block   *walk.StructDefine
}

type SrvDefine struct {
	Package string
	Name    string
	Req     *walk.StructDefine
	Rsp     *walk.StructDefine
}

type pbDefine struct {
	Name  string
	Field []pbField
}

type pbField struct {
	Type string
	Name string
}
