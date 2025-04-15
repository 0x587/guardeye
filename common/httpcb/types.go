package httpcb

type BasicCallbackConfig struct {
	Online  CallbackItem `json:"online"`
	Offline CallbackItem `json:"offline"`
}

type CallbackConfig struct {
	Basic BasicCallbackConfig `json:"basic"`
	Data  []DataCallbackItem  `json:"data"`
}

type CallbackItem struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Enabled bool              `json:"enabled"`
}

type DataCallbackItem struct {
	Id           string            `json:"id"`
	Url          string            `json:"url"`
	Method       string            `json:"method"`
	Headers      map[string]string `json:"headers"`
	Enabled      bool              `json:"enabled"`
	TopicPattern string            `json:"topicPattern"`
}
