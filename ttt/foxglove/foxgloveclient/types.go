package foxgloveclient

type channelInfo struct {
	Topic      string
	SchemaName string
	Schema     string
}

type message struct {
	Topic      string
	SchemaName string
	JsonData   string
}
