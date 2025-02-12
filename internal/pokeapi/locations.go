// Concerns pokedex api location data requests and responses 

package pokeapi

import (
	"encoding/json"
	"fmt"
)

type LocationAreaResponse struct {
        Count           int             `json:"count"`
        Next            *string         `json:"next"`
        Previous        *string         `json:"previous"`
        Results         []struct {
                Name    string `json:"name"`
                URL     string `json:"url"`
        } `json:"results"`
}

func (c *Client) GetLocationAreas(url string) (LocationAreaResponse, error) {

	// Check for cached response first
	if cachedVal, exists := c.cache.Get(url); exists {
		var cachedResponse LocationAreaResponse
		if err := json.Unmarshal(cachedVal, &cachedResponse); err != nil {
			return LocationAreaResponse{}, fmt.Errorf("Error unmarshalling cached response: %w", err)
		}
		return cachedResponse, nil
	}

	// If not in cache make request to api
	res, err := c.httpClient.Get(url)
	if err != nil { 
		return LocationAreaResponse{}, fmt.Errorf("error creating request: %w", err)

	}
	defer res.Body.Close()

	var locationAreaResponse LocationAreaResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationAreaResponse); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("error decoding response body: %w", err)
	}

	// Add response to cache
	val, err := json.Marshal(locationAreaResponse)
	if err != nil{
		return locationAreaResponse, fmt.Errorf("error marshalling response into cache: %w", err)
	}
	c.cache.Add(url, val) 

	return locationAreaResponse, nil
}
