package main

import (
	"strings"
)

func CleanUpMessage(message string) string {
	// Remove stars
	message = strings.ReplaceAll(message, "*", "")

	// Remove leading and trailing whitespaces
	message = strings.TrimSpace(message)

	return message
}

func main() {
	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************`

	cleanedMsg := CleanUpMessage(message)
	println(cleanedMsg)
}
