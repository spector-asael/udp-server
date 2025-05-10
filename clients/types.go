package clients

type clientMessage struct {
	ID       string
	Username string
	IP       string
	Port     string
	Message  string // Updated to uppercase to match the server
	Type     string // "ping" or "content"
}
