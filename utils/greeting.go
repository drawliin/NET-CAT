package utils

import (
	"os"
)

func Greeting() string {

	//read greeting from greeting.txt
	bytes, err := os.ReadFile("greeting.txt")

	if err != nil {
		greeting := "Welcome to TCP-Chat!"
		return greeting
	}
	return string(bytes)

}
