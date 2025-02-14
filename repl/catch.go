package repl

import (
	"fmt"
	"math/rand"
)


func Catch (cfg *Config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("please provide a pokemon to catch")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name) 

	// Api call
	pokemon, err := cfg.PokeClient.GetPokemon(name)
	if err != nil {
		return fmt.Errorf("there was an error retrieving %s's data: %w\n", name, err)
	}

	// Determine if pokemon is caught
	prob := calculateCatchProbability(pokemon.BaseExp)
	rand := rand.Float64()
	fmt.Println()
	fmt.Println(prob)
	fmt.Println(rand)
	isCaught := rand < prob
	fmt.Println(isCaught)
	fmt.Println()
	
	if !isCaught {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	
	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.Collection.Add(pokemon)
	return nil
}

func calculateCatchProbability(baseExp int) float64 {
	if baseExp < 1 {
		baseExp = 1
	} 
	if baseExp > 350 {
		baseExp = 350
	}

	baseProb := 0.9 - (float64(baseExp) / 350.0 * 0.8)
	randomVariation := 0.05 * (rand.Float64()*2-1)
	return baseProb + randomVariation
}
