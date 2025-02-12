// Concerns pokedex api location data requests and responses 

package pokeapi

import (
	"encoding/json"
	"fmt"
)

const (
	baseURL = "https://pokeapi.co/api/v2/"
)

type LocationsResp struct {
        Count           int             `json:"count"`
        Next            *string         `json:"next"`
        Previous        *string         `json:"previous"`
        Results         []struct {
                Name    string 	`json:"name"`
                URL     string 	`json:"url"`
        } `json:"results"`
}

type LocationDetailsResp struct {
	Name		string	`json:"name"`
	Encounters 	[]struct {
		Pokemon	Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name	string	`json:"name"`
	URL	string	`json:"url"`
} 

func (c *Client) GetLocations(pageURL *string) (LocationsResp, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check for cached response first
	if cachedVal, exists := c.cache.Get(url); exists {
		var cachedResponse LocationsResp
		if err := json.Unmarshal(cachedVal, &cachedResponse); err != nil {
			return LocationsResp{}, fmt.Errorf("Error unmarshalling cached response: %w", err)
		}
		return cachedResponse, nil
	}

	// If not in cache make request to api
	res, err := c.httpClient.Get(url)
	if err != nil { 
		return LocationsResp{}, fmt.Errorf("error creating request: %w", err)

	}
	defer res.Body.Close()

	var locationsResp LocationsResp
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationsResp); err != nil {
		return LocationsResp{}, fmt.Errorf("error decoding response body: %w", err)
	}

	// Add response to cache
	val, err := json.Marshal(locationsResp)
	if err != nil{
		return locationsResp, fmt.Errorf("error marshalling response into cache: %w", err)
	}
	c.cache.Add(url, val) 

	return locationsResp, nil
}

func (c *Client) GetLocationDetails(location string) (LocationDetailsResp, error) {
	url := baseURL + "location-area/" + location

	// Check for cached response first
	if cachedVal, exists := c.cache.Get(url); exists {
		var cachedResponse LocationDetailsResp
		if err := json.Unmarshal(cachedVal, &cachedResponse); err != nil {
			return LocationDetailsResp{}, fmt.Errorf("Error unmarshalling cached response: %w", err)
		}
		return cachedResponse, nil
	}

	// If not in cache make request to api
	res, err := c.httpClient.Get(url)
	if err != nil { 
		return LocationDetailsResp{}, fmt.Errorf("error creating request: %w", err)

	}
	defer res.Body.Close()

	var locationDetails LocationDetailsResp
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locationDetails); err != nil {
		return LocationDetailsResp{}, fmt.Errorf("error decoding response body: %w", err)
	}

	// Add response to cache
	val, err := json.Marshal(locationDetails)
	if err != nil{
		return locationDetails, fmt.Errorf("error marshalling response into cache: %w", err)
	}
	c.cache.Add(url, val) 

	return locationDetails, nil

}
