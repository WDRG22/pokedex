// Concerns handling user commands 

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

// Get next locations. Update cfg w/ new URLs 
func Map(cfg *Config, args []string) error {
	if cfg.NextURL == nil{
		fmt.Println("you're on the last page")
		return nil	
	}

	// Get locations 
	locationsResp, err := cfg.PokeClient.GetLocations(cfg.NextURL)
	if err != nil {
		return err
	}

	// Update config
	cfg.PrevURL = cfg.NextURL
	if locationsResp.Next == nil {
		cfg.NextURL = nil 
	} else {
		cfg.NextURL = locationsResp.Next
	}
	
	// Print locations 
	locations := locationsResp.Results
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

// Get previous locations. Update cfg w/ new URLs 
func MapB(cfg *Config, args []string) error {
	if cfg.PrevURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	// Get locations 
	locationsResp, err := cfg.PokeClient.GetLocations(cfg.PrevURL)
	if err != nil {
		return err
	}

	// Update config
	cfg.NextURL = cfg.PrevURL
	if locationsResp.Previous == nil {
		cfg.PrevURL = nil
	} else {
		cfg.PrevURL = locationsResp.Previous
	}

	// Print locations 
	locations := locationsResp.Results
	for _, location := range locations{
		fmt.Println(location.Name)
	}
	return nil
}

// Get pokemon in this location 
func Explore(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a location")
	}

	// Api call
	locationDetails, err:= cfg.PokeClient.GetLocationDetails(args[0])
	if err != nil {
		return fmt.Errorf("there was an error retrieving this location's pokemon: %w", err)
	}

	for _, encounter := range locationDetails.Encounters{
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
