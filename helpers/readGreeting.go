package helpers

import "os"

func ReadGreeting() []byte {
	defaultGreeting := "Welcome to TCP-Chat!\n"
	logoBytes, err := os.ReadFile("greeting.txt")
	if err != nil {
		return []byte(defaultGreeting)
	}
	return logoBytes
}
