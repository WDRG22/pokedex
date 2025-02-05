package commands

import (
	"fmt"
	"os"
)

type cliCommand struct {
        name            string
        description     string
        callback        func() error
}

var commandRegistry map[string]cliCommand 

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Useage:")
	for _, command := range commandRegistry {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}


func int() {
	commandRegistry = map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exits the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
}
