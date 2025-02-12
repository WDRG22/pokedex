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

	return locationAreaResponse, nil
}
