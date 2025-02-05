package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	commands "github.com/wdrg22/pokedex/commands" 
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		if command, ok := commands.CommandRegistry[cleanedInput[0]]; ok {
			err := command.Callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}		
}

func cleanInput(text string) []string {
	cleaned := strings.Fields(text)
	for i, str := range cleaned {
		cleaned[i] = strings.ToLower(strings.Trim(str, "!?.,- "))
	}

	return cleaned
}
