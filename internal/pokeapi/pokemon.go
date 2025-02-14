package pokeapi

import (
	"fmt"
	"encoding/json"
	"github.com/wdrg22/pokedex/internal/pokedex"
)

type PokemonResp struct {
	Name	string 	`json:"name"`
	BaseExp	int	`json:"base_experience"`	
}

// Convert API response to game Pokemon type
func (p *PokemonResp) ToPokemon() pokedex.Pokemon {
	return pokedex.Pokemon{
		Name: 		p.Name,
		BaseExp: 	p.BaseExp,
	}
}

func (c *Client) GetPokemon(name string) (pokedex.Pokemon, error) {
	url := baseURL + "pokemon/" + name

	// Check for cached response
	if cachedVal, exists := c.cache.Get(url); exists {
		var cachedResponse PokemonResp
		if err := json.Unmarshal(cachedVal, &cachedResponse); err != nil {
			return pokedex.Pokemon{}, fmt.Errorf("error unmarshalling cached response: %w\n", err)
		}
		return cachedResponse.ToPokemon(), nil
	}

	// If not in cache make api request
	res, err := c.httpClient.Get(url)
	if err != nil {
		return pokedex.Pokemon{}, fmt.Errorf("error creating request: %w\n", err)
	}
	defer res.Body.Close()

	var pokemonResp PokemonResp
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&pokemonResp); err != nil {
		return pokedex.Pokemon{}, fmt.Errorf("error decoding response body: %w\n", err)
	}

	// Add response to cache
	val, err := json.Marshal(pokemonResp)
	if err != nil {
		return pokemonResp.ToPokemon(), fmt.Errorf("error marshalling response into cache: %w\n", err)
	}
	c.cache.Add(url, val)
	
	// Return the pokemon as a pokedex.Pokemon type
	return pokemonResp.ToPokemon(), nil
}
