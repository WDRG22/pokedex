package repl

import (
	"fmt"
	"sort"
)

func Pokedex(cfg *Config, args []string) error {
	// Check if collection empty
	if len (cfg.Collection.Caught) == 0 {
		return fmt.Errorf("you don't have any pokemon in your collection yet!")
	}

	pokemon := []string{}
	for key, _ := range cfg.Collection.Caught {
		pokemon = append(pokemon, key)
	}


	sort.Strings(pokemon)
	fmt.Println("Your Pokedex:")
	for _, name := range pokemon {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}
