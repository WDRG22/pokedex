// Concerns api interactions

package pokeapi

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type LocationAreaResponse struct {
	Count		int 		`json:"count"`
	Next 		*string 	`json:"next"`
	Previous 	*string 	`json:"previous"`
	Results 	[]LocationArea 	`json:"results"`
}

type LocationArea struct {
	Name 	string `json:"name"`
	URL 	string `json:"url"`
}


func GetLocationAreas(url string) (LocationAreaResponse, error) {
	res, err := http.Get(url)
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
	
