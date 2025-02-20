package ws

//type PubSub interface {
//	ServeWs(w http.ResponseWriter, r *http.Request)
//	Broadcast(key string, p []byte)
//}
//
//func NewPubSub() PubSub {
//	res := &pubSub{
//		listener: make(map[string][]uuid.UUID),
//	}
//	res.IF = New(res.handler)
//	return res
//}
//
//type pubSub struct {
//	IF
//	listener map[string][]uuid.UUID
//}
//
//func (ps *pubSub) Broadcast(key string, p []byte) {
//	fmt.Println("bc ", ps.listener[key], key, string(p))
//	for _, cid := range ps.listener[key] {
//		ps.IF.Send(cid, p)
//	}
//}
//
//func (ps *pubSub) handler(h *hub, ms *msg) {
//	m := &types.WsMsgBase{}
//	if err := json.Unmarshal(ms.msg, m); err != nil {
//		return
//	}
//	switch m.Cmd {
//	case "Sub":
//		req := &types.WsSubReq{}
//		if err := json.Unmarshal(ms.msg, req); err != nil {
//			return
//		}
//		ps.listener[req.SubKey] = append(ps.listener[req.SubKey], ms.cid)
//	}
//}
