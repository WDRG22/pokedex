// Concerns location data request and response methods

package pokeapi

import (
	"encoding/json"
	"fmt"
)


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
