package commands


type Config struct {
        NextURL string
        PrevURL string
}

func InitConfig(nextURL string, prevURL string) *Config {
	return &Config{nextURL, prevURL}
}
