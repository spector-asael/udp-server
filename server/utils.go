package server

import (
	"fmt"
	"net"
	"time"
)

func handleClientMessage(conn *net.UDPConn, message clientMessage, clientAddr *net.UDPAddr) {
	// Check if the client ID exists in the map
	if _, exists := connectedClients[message.ID]; !exists {
		// Add new client to the map
		connectedClients[message.ID] = clientAddr
		fmt.Printf("New client connected: %s (%s:%s)\n", message.Username, message.IP, message.Port)

		// Send connection success message
		response := "Connection successful!"
		_, err := conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
		return
	}

	// Handle message based on its type
	switch message.Type {
	case "ping":
		// Respond to ping
		response := "Ping received. Connection is active."
		_, err := conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Println("Error sending ping response:", err)
		}

	case "content":
		// Wait for 1-2 seconds to ensure all data is received
		time.Sleep(2 * time.Second)

		// Broadcast the message to all connected clients
		for id, addr := range connectedClients {
			if id != message.ID { // Don't send the message back to the sender
				broadcast := fmt.Sprintf("%s: %s", message.Username, message.Message)
				_, err := conn.WriteToUDP([]byte(broadcast), addr)
				if err != nil {
					fmt.Printf("Error broadcasting to %s: %v\n", addr.String(), err)
				}
			}
		}

	default:
		fmt.Println("Unknown message type:", message.Type)
	}
}
