package server

// Define a struct to match the client's message
type clientMessage struct {
	ID       string
	Username string
	IP       string
	Port     string
	Message  string // Note: Capitalized field name to match JSON unmarshaling
	Type     string // "ping" or "content"
}
