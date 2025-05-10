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

	// Example send message
	message := clientMessage{
		ID:       userID,
		Username: username,
		IP:       fmt.Sprintf("%v", addr),
		Port:     port,
		message:  "Hello from client!",
	}

	data, err := json.Marshal(message)
	if err != nil {
		fmt.Println("Failed to marshal message:", err)
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Failed to send message:", err)
	}
}
