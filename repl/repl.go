package repl

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	commands "github.com/wdrg22/pokedex/commands"
)

func StartREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := CleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		if cmd, ok := commands.CommandRegistry[cleanedInput[0]]; ok {
			err := cmd.Callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}		
}

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
