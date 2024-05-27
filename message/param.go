package message

const (
	ProtocolVersion uint32 = 70015
	Services        uint64 = 0
	UserAgent       string = "/btc-handshake:0.1.0/"
	CommandVersion  string = "version"
	CommandVerack   string = "verack"
	DefaultPort     string = "8333"
)
