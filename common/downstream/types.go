package downstream

type CommandReq struct {
	Id     string         `json:"id"`
	Action string         `json:"action"`
	Data   CommandReqData `json:"data"`
}

type CommandReqData struct {
	RosTopic string `json:"ros_topic"`
	Data     any    `json:"data"`
}

type CommandRsp struct {
	Id   string `json:"id"`
	Ok   bool   `json:"ok"`
	Data any    `json:"data"`
}
