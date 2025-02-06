// Concerns user input and the Read-Eval-Print Loop

package repl

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	commands "github.com/wdrg22/pokedex/commands"
)

func StartREPL(cfg *commands.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := CleanInput(input)

		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]

		if cmd, exists := commands.CommandRegistry[commandName]; exists {
			err := cmd.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}		
}

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
