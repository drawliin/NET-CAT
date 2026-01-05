package helpers

import "os"

func GetGreeting() string {
	greeting := "Welcome to TCP-Chat!\n[ENTER YOUR NAME]:"
	bytes, err := os.ReadFile("linux_logo.txt")
	if err != nil {
		return greeting
	}
	greeting = string(bytes)
	return greeting
}
