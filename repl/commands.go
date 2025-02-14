package repl 

import (
	"fmt"
	"os"
	"sort"
)


type cliCommand struct {
        name            string
        description     string
        callback        func(*Config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exits the Pokedex",
			callback: Exit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: Help,
		},
		"map": {
			name: "map",
			description: "Enumerates the next map locations",
			callback: Map,
		},
		"mapb": {
			name: "mapb",
			description: "Enumerates the previous map locations",
			callback: MapB,
		},
		"explore": {
			name: "explore",
			description: "Displays a list of pokemon in a location",
			callback: Explore,
		},
		"catch": {
			name: "catch",
			description: "Attempts to catch a pokemon",
			callback: Catch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspects a pokemon from your collection",
			callback: Inspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Lists all the pokemon in your collection",
			callback: Pokedex,
		},
	}
}

func Exit(cfg *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func Help(cfg *Config, args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Useage:")
	fmt.Println()


	commands := getCommands()
	commandNames := []string{}
	for name := range commands{
		commandNames = append(commandNames, name)
	}
	sort.Strings(commandNames)
	
	for _, name := range commandNames{
		command := commands[name]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
