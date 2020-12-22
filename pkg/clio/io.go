package clio

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Get gets text input from stdin.
func Get(instructions string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(instructions + ": ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input\n%w", err)
	}

	return strings.ReplaceAll(text, "\n", ""), nil
}

// MustGet get text input from stdin and exits on error.
func MustGet(instructions string) string {
	text, err := Get(instructions)
	if err != nil {
		log.Fatal(err)
	}

	return text
}
