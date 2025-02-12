// Concerns handling user commands 

package repl 

import (
	"fmt"
	"os"
	"sort"
	"github.com/wdrg22/pokedex/internal/pokeapi"
)


type cliCommand struct {
        name            string
        description     string
        callback        func(*Config) error
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
			description: "Enumerates the next map location areas",
			callback: Map,
		},
		"mapb": {
			name: "mapb",
			description: "Enumerates the previous map location areas",
			callback: MapB,
		},
	}
}

func Exit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func Help(cfg *Config) error {
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

// Get next location areas. Update cfg w/ new URLs 
func Map(cfg *Config) error {
	if cfg.NextURL == ""{
		fmt.Println("you're on the last page")
		return nil	
	}

	// Get location areas
	locationAreaResponse, err := pokeapi.GetLocationAreas(cfg.NextURL)
	if err != nil {
		return err
	}

	// Update config
	cfg.PrevURL = cfg.NextURL
	if locationAreaResponse.Next == nil {
		cfg.NextURL = ""
	} else {
		cfg.NextURL = *locationAreaResponse.Next
	}
	
	// Print location areas
	locationAreas := locationAreaResponse.Results
	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}
	return nil
}

// Get previous location areas. Update cfg w/ new URLs 
func MapB(cfg *Config) error {
	if cfg.PrevURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	// Get location areas
	locationAreaResponse, err := pokeapi.GetLocationAreas(cfg.PrevURL)
	if err != nil {
		return err
	}

	// Update config
	cfg.NextURL = cfg.PrevURL
	if locationAreaResponse.Previous == nil {
		cfg.PrevURL = ""
	} else {
		cfg.PrevURL = *locationAreaResponse.Previous
	}

	// Print location areas
	locationAreas := locationAreaResponse.Results
	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}
	return nil
}
