package config

type Config struct {
	Port string `json:"port"`
}

func LoadConfigs() Config {
	return Config{
		Port: "3000",
	}
}
