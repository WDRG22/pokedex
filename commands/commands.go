package commands

import (
	"fmt"
	"os"
)

type CliCommand struct {
        Name            string
        Description     string
        Callback        func() error
}

var CommandRegistry map[string]CliCommand 

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Useage:")
	for _, command := range CommandRegistry {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}


func init() {
	CommandRegistry = map[string]CliCommand{
		"exit": {
			Name: "exit",
			Description: "Exits the Pokedex",
			Callback: commandExit,
		},
		"help": {
			Name: "help",
			Description: "Displays a help message",
			Callback: commandHelp,
		},
	}
}
