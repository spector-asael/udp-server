package server

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

var (
	connectedClients = make(map[string]*net.UDPAddr)
	lastSeen         = make(map[string]time.Time)
	mu               sync.Mutex
)

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

	go monitorInactiveClients()

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
