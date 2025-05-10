package server

import (
	"encoding/json"
	"fmt"
	"net"
)

// Define a struct to match the client's message
type clientMessage struct {
	ID       string
	Username string
	IP       string
	Port     string
	Message  string // Note: Capitalized field name to match JSON unmarshaling
	Type     string // "ping" or "content"
}

var connectedClients = make(map[string]*net.UDPAddr)

func StartServer() {
	var port string
	fmt.Print("Enter the port: ")
	fmt.Scanln(&port)

	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error starting UDP server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Server started on port", port)

	buf := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		// Decode the received JSON data
		var message clientMessage
		err = json.Unmarshal(buf[:n], &message)
		if err != nil {
			fmt.Println("Failed to decode message:", err)
			continue
		}

		// Pass conn to handleClientMessage
		handleClientMessage(conn, message, clientAddr)
	}
}
