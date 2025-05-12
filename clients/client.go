package clients

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
)

func StartClient() {

	username := generateDefaultUsername()
	ID := generateRandomID()

	for {
		fmt.Printf("Welcome, %s! what would you like to do?\n", username)
		fmt.Println("1. Connect to server")
		fmt.Println("2. Update username")
		fmt.Println("3. Exit")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			initiateConnection(username, ID)

		case "2":
			username = updateUsername()

		case "3":
			fmt.Println("Exiting client menu...")
			os.Exit(1)

		default:
			fmt.Println("Invalid option. Please enter 1, 2, or 3.")
		}
	}
}

func generateRandomID() string {

	return fmt.Sprintf("%d", rand.Intn(1000000))
}

func initiateConnection(username string, userID string) {
	var port string
	fmt.Print("Enter the server port to connect to (e.g., 8080): ")
	fmt.Scanln(&port)

	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+port)
	if err != nil {
		fmt.Println("Invalid port or address:", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Failed to connect:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server at", addr.String())

	// Send an initial connection message
	initialMessage := clientMessage{
		ID:       userID,
		Username: username,
		IP:       fmt.Sprintf("%v", addr),
		Port:     port,
		Type:     "ping",               // Initial message type
		Message:  "Hello from client!", // Updated field name
	}

	data, err := json.Marshal(initialMessage)
	if err != nil {
		fmt.Println("Failed to marshal initial message:", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Failed to send initial message:", err)
		return
	}

	go startHeartbeat(conn, addr, userID, username)
	// Start the message loop
	startMessageLoop(conn, username, userID, addr, port)
}

func startMessageLoop(conn *net.UDPConn, username, userID string, addr *net.UDPAddr, port string) {
	// Start a goroutine to listen for incoming messages
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, _, err := conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println("Error reading from server:", err)
				return
			}
			fmt.Println(string(buffer[:n]))
		}
	}()

	// Print the prompt once
	fmt.Println("Enter your message (or type 'exit' to quit):")

	// Main loop to send messages
	for {
		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		message := clientMessage{
			ID:       userID,
			Username: username,
			IP:       fmt.Sprintf("%v", addr),
			Port:     port,
			Type:     "content", // Message type
			Message:  input,     // Updated field name
		}

		data, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Failed to marshal message:", err)
			continue
		}

		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Failed to send message:", err)
		}
	}
}
