package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// Starts the UDP server and listens for incoming messages
func StartServer(port string) {
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to resolve address: %v", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on UDP: %v", err)
	}
	defer conn.Close()

	server := NewServer()

	buffer := make([]byte, 1024)

	fmt.Println("UDP Chat Server started on port", port)

	go server.CleanupInactiveClients()

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading from UDP: %v", err)
			continue
		}

		var msg Message
		if err := json.Unmarshal(buffer[:n], &msg); err != nil {
			log.Printf("Invalid message format from %v: %v", clientAddr, err)
			continue
		}

		// Register new client
		server.RegisterClient(msg.ID, msg.Username, clientAddr)

		// Refresh client's last-seen time
		server.UpdateHeartbeat(msg.ID)

		// Send acknowledgment
		ack := []byte("ACK")
		conn.WriteToUDP(ack, clientAddr)

		// Handle message type
		if strings.HasPrefix(msg.Content, "/") {
			server.HandleCommand(msg, conn)
		} else {
			server.BroadcastMessage(msg, conn)
		}
	}
}
package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// Starts the UDP server and listens for incoming messages
func StartServer(port string) {
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to resolve address: %v", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on UDP: %v", err)
	}
	defer conn.Close()

	server := NewServer()

	buffer := make([]byte, 1024)

	fmt.Println("UDP Chat Server started on port", port)

	go server.CleanupInactiveClients()

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading from UDP: %v", err)
			continue
		}

		var msg Message
		if err := json.Unmarshal(buffer[:n], &msg); err != nil {
			log.Printf("Invalid message format from %v: %v", clientAddr, err)
			continue
		}

		// Register new client
		server.RegisterClient(msg.ID, msg.Username, clientAddr)

		// Refresh client's last-seen time
		server.UpdateHeartbeat(msg.ID)

		// Send acknowledgment
		ack := []byte("ACK")
		conn.WriteToUDP(ack, clientAddr)

		// Handle message type
		if strings.HasPrefix(msg.Content, "/") {
			server.HandleCommand(msg, conn)
		} else {
			server.BroadcastMessage(msg, conn)
		}
	}
}
