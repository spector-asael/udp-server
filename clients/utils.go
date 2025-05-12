package clients

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func generateDefaultUsername() string {

	var adjectives = []string{
		"Happy", "Brave", "Clever", "Mighty", "Sneaky",
		"Zany", "Jolly", "Witty", "Curious", "Gentle",
	}

	var nouns = []string{
		"Tiger", "Falcon", "Wizard", "Pirate", "Ninja",
		"Robot", "Unicorn", "Panther", "Dragon", "Knight",
	}

	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	number := strconv.Itoa(rand.Intn(100))

	return fmt.Sprint(adjective + noun + number)
}

func updateUsername() string {
	scanner := bufio.NewScanner(os.Stdin)
	validUsername := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

	for {
		fmt.Print("Enter a new username (letters, numbers, and underscores only): ")
		scanner.Scan()
		newUsername := strings.TrimSpace(scanner.Text())

		if !validUsername.MatchString(newUsername) {
			fmt.Println("Invalid username. Please use only letters, numbers, and underscores.")
			continue
		}

		return newUsername
	}
}

func startHeartbeat(conn *net.UDPConn, addr *net.UDPAddr, id, username string) {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			msg := clientMessage{
				ID:       id,
				Username: username,
				IP:       addr.IP.String(),
				Port:     fmt.Sprint(addr.Port),
				Type:     "ping",
			}
			data, _ := json.Marshal(msg)
			conn.WriteToUDP(data, addr)
		}
	}()
}
