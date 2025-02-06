package pokeapi

type Config struct {
	NextURL string
	PrevURL string
}

type LocationArea {
	Name: string
	URL: string
}

func GetLocationAreas(config *Config) []string {
	res, err := http.Get(config.NextURL)
	if err != nil { 
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code %d and\nbody: %s\n", res.StatusCode, body) 
	}

	return []string{"test","getlocationareas","return"}
}
	
