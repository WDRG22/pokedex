package main

import (
	"time"
	"github.com/wdrg22/pokedex/repl"
	"github.com/wdrg22/pokedex/internal/pokeapi"
)


func main() {
	client := pokeapi.NewClient(5 * time.Second)
	config := repl.InitConfig(client, "https://pokeapi.co/api/v2/location-area/", "")
	repl.StartREPL(config)
}
