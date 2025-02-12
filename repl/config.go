package repl

import "github.com/wdrg22/pokedex/internal/pokeapi"

type Config struct {
        Client  *pokeapi.Client
        NextURL string
        PrevURL string
}

func InitConfig(httpClient *pokeapi.Client, next string, prev string) *Config {
	return &Config{
		Client: httpClient,
		NextURL: next,
		PrevURL: prev,
	}
}
