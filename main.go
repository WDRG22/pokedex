package main

import (
	"time"
	"github.com/wdrg22/pokedex/repl"
	"github.com/wdrg22/pokedex/internal/pokeapi"
	"github.com/wdrg22/pokedex/internal/pokedex"
)


func main() {
	client := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	collection := pokedex.NewCollection()
	config := repl.InitConfig(client, collection, "https://pokeapi.co/api/v2/location-area/", "")
	repl.StartREPL(config)
}
