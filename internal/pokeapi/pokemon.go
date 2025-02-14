package pokeapi

import (
	"fmt"
	"encoding/json"
	"github.com/wdrg22/pokedex/internal/pokedex"
)

type PokemonResp struct {
	Name	string 	`json:"name"`
	BaseExp	int	`json:"base_experience"`	
	Height	int 	`json:"height"`
	Weight	int	`json:"weight"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat struct{
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct{
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}


// Convert Pokemon API response to game pokemon type
func (p *PokemonResp) ToPokemon() pokedex.Pokemon {

	// Convert Stats
	stats := make([]pokedex.Stat, len(p.Stats))
	for i, s := range p.Stats {
		stats[i] = pokedex.Stat{
			Name:     s.Stat.Name,
			BaseStat: s.BaseStat,
		}
	}

	// Convert Types
	types := make([]pokedex.Type, len(p.Types))
	for i, t := range p.Types {
		types[i] = pokedex.Type{
			Name: t.Type.Name,
		}
	}

	return pokedex.Pokemon{
		Name:    p.Name,
		BaseExp: p.BaseExp,
		Height:  p.Height,
		Weight:  p.Weight,
		Stats:   stats,
		Types:   types,
	}
}

// Get individual pokemon data from api
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
