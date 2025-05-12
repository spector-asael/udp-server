package server

import (
	"net"
	"testing"
)

func BenchmarkHandleClientMessage(b *testing.B) {
	// Setup a dummy UDP address and connection
	addr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9999}
	conn, err := net.ListenUDP("udp", nil) // nil assigns a random available port
	if err != nil {
		b.Fatalf("Failed to listen on UDP: %v", err)
	}
	defer conn.Close()

	msg := clientMessage{
		ID:       "abc123",
		Username: "TestUser",
		IP:       "127.0.0.1",
		Port:     "9999",
		Type:     "ping",
		Message:  "Hello!",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handleClientMessage(conn, msg, addr)
	}
}
