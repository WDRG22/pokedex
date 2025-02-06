package main

import (
	repl "github.com/wdrg22/pokedex/repl"
	commands "github.com/wdrg22/pokedex/commands"
)


func main() {
	config := commands.InitConfig("https://pokeapi.co/api/v2/location-area/", "")
	repl.StartREPL(config)
}

