package main

import (
	"fmt"
	"os"
	"udp-chat-app/clients"
	"udp-chat-app/server"
)

func main() {

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1) Start Server")
		fmt.Println("2) Start Client")
		fmt.Println("3) Exit")
		fmt.Println(" ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			// handle starting the server

			server.StartServer()

		case "2":
			// Handle starting the client
			clients.StartClient()

		case "3":
			// Exit the program
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			// Handle invalid input
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		}
	}

}
