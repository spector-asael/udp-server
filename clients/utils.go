package clients

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
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
