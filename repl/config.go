package repl

import (
	"github.com/wdrg22/pokedex/internal/pokeapi"
	"github.com/wdrg22/pokedex/internal/pokedex"
)

type Config struct {
        PokeClient  	*pokeapi.Client
	Collection	*pokedex.Collection
        NextURL 	*string
        PrevURL 	*string
}


func InitConfig(httpClient *pokeapi.Client, collection *pokedex.Collection, next string, prev string) *Config {
	return &Config{
		PokeClient: 	httpClient,
		Collection:	collection,
		NextURL: 	&next,
		PrevURL: 	&prev,
	}
}
