package repl

import (
	"fmt"
)

// Get next locations. Update cfg w/ new URLs 
func Map(cfg *Config, args []string) error {
	if cfg.NextURL == nil{
		fmt.Println("you're on the last page")
		return nil	
	}

	// Api call 
	locationsResp, err := cfg.PokeClient.GetLocations(cfg.NextURL)
	if err != nil {
		return fmt.Errorf("there was an error retrieving these locations %w\n", err)
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

	// Api call 
	locationsResp, err := cfg.PokeClient.GetLocations(cfg.PrevURL)
	if err != nil {
		return fmt.Errorf("there was an error retrieving these locations: %w\n", err)
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


