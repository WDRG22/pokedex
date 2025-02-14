package repl

import (
	"fmt"
)

func Inspect(cfg *Config, args []string) error{
	if len(args) < 1 {
		return fmt.Errorf("please provide a pokemon to inspect")
	}

	name := args[0]

	// Check if pokemon in user's collection
	pokemon, exists := cfg.Collection.Caught[name]
	if !exists {
		return fmt.Errorf("%s does not appear to be in your collection", name)
	}
	
	// Print pokemon's data
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n",stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types{
		fmt.Printf("  -%s\n", t.Name)
	}

	return nil
}
