package ws

type MsgBase struct {
	Cmd string `json:"cmd"`
}

type InitRsp struct {
	MsgBase
	SessionId string `json:"sessionId"`
}
type MetricRsp struct {
	MsgBase
	Name  string `json:"name"`
	Value string `json:"value"`
}
