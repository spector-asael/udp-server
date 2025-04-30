package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Connect to chat")
		fmt.Println("2. Change username")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")

		if !scanner.Scan() {
			fmt.Println("Error reading input.")
			break
		}

		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			// Connect to chat (implementation to be added)
		case "2":
			// Change username (implementation to be added)
		case "3":
			// Exit (implementation to be added)
		default:
			fmt.Println("Invalid option. Please enter 1, 2, or 3.")
			fmt.Println("")
		}
	}
}
