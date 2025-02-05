package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(cleanInput("?tesT!, Data .wOrd, bum-"))
}

func cleanInput(text string) []string {
	cleaned := strings.Fields(text)
	for i, str := range cleaned {
		cleaned[i] = strings.ToLower(strings.Trim(str, "!?.,- "))
	}

	return cleaned
}
