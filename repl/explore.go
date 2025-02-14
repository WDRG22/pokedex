package repl 

import (
	"fmt"
)

// Get pokemon in this location 
func Explore(cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a location")
	}
	location := args[0]

	// Api call
	locationDetails, err:= cfg.PokeClient.GetLocationDetails(location)
	if err != nil {
		return fmt.Errorf("there was an error retrieving this location's pokemon: %w\n", err)
	}

	for _, encounter := range locationDetails.Encounters{
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
