package ros

type FieldDefine struct {
	Name        string
	Type        string
	IsArray     bool
	IsComplex   bool
	ComplexType *MsgDefine
}

type BlockDefine struct {
	Fields []*FieldDefine
}

type MsgDefine struct {
	Package string
	Name    string
	Block   *BlockDefine
}

type SrvDefine struct {
	Package string
	Name    string
	Req     *BlockDefine
	Rsp     *BlockDefine
}

type Message struct {
	Topic string
	Msg   *MsgDefine
}

type Service struct {
	Topic string
	Srv   *SrvDefine
}
