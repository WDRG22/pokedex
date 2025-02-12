package repl

import "github.com/wdrg22/pokedex/internal/pokeapi"

type Config struct {
        PokeClient  *pokeapi.Client
        NextURL string
        PrevURL string
}

func InitConfig(httpClient *pokeapi.Client, next string, prev string) *Config {
	return &Config{
		PokeClient: httpClient,
		NextURL: next,
		PrevURL: prev,
	}
}
