package foxgloveclient

type channelInfo struct {
	Id         int    `json:"id"`
	Topic      string `json:"topic"`
	Encoding   string `json:"encoding"`
	SchemaName string `json:"schemaName"`
	Schema     string `json:"schema"`
}

type serviceInfo struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	RequestSchema  string `json:"requestSchema,omitempty"`
	ResponseSchema string `json:"responseSchema,omitempty"`
	Request        struct {
		Encoding       string `json:"encoding"`
		SchemaName     string `json:"schemaName"`
		SchemaEncoding string `json:"schemaEncoding"`
		Schema         string `json:"schema"`
	} `json:"request,omitempty"`
	Response struct {
		Encoding       string `json:"encoding"`
		SchemaName     string `json:"schemaName"`
		SchemaEncoding string `json:"schemaEncoding"`
		Schema         string `json:"schema"`
	} `json:"response,omitempty"`
}

type message struct {
	Topic      string
	SchemaName string
	TransData  string
}
