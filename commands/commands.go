// Concerns user commands handling and state management

package commands

import (
	"fmt"
	"os"
	pokeapi "github.com/wdrg22/pokedex/internal"
)


type Config struct {
	NextURL string
	PrevURL string
}

type CliCommand struct {
        Name            string
        Description     string
        Callback        func(*Config) error
}

var CommandRegistry map[string]CliCommand 

func InitConfig(nextURL string, prevURL string) *Config {
	return &Config{nextURL, prevURL}

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
		"map": {
			Name: "map",
			Description: "Enumerates the next map location areas",
			Callback: commandMap,
		},
		"mapb": {
			Name: "mapb",
			Description: "Enumerates the previous map location areas",
			Callback: commandMapB,
		},
	}
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Useage:")
	fmt.Println()
	for _, command := range CommandRegistry {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	fmt.Println()
	return nil
}

// Get next location areas. Update config w/ new URLs 
func commandMap(config *Config) error {
	if config.NextURL == ""{
		fmt.Println("you're on the last page")
		return nil	
	}

	// Get location areas
	locationAreaResponse, err := pokeapi.GetLocationAreas(config.NextURL)
	if err != nil {
		return err
	}

	// Update config
	config.PrevURL = config.NextURL
	if locationAreaResponse.Next == nil {
		config.NextURL = ""
	} else {
		config.NextURL = *locationAreaResponse.Next
	}
	
	// Print location areas
	locationAreas := locationAreaResponse.Results
	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}
	return nil
}

// Get previous location areas. Update config w/ new URLs 
func commandMapB(config *Config) error {
	if config.PrevURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	// Get location areas
	locationAreaResponse, err := pokeapi.GetLocationAreas(config.PrevURL)
	if err != nil {
		return err
	}

	// Update config
	config.NextURL = config.PrevURL
	if locationAreaResponse.Previous == nil {
		config.PrevURL = ""
	} else {
		config.PrevURL = *locationAreaResponse.Previous
	}

	// Print location areas
	locationAreas := locationAreaResponse.Results
	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}
	return nil
}
