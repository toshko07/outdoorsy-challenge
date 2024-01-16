package configs

import (
	"github.com/labstack/gommon/log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env        string `required:"true"`
	ServerPort int    `required:"true" split_words:"true"`
	DB         DB     `required:"true"`
}

func readConfig(filename string) (*Config, error) {
	if err := godotenv.Load(filename); err != nil {
		log.Fatalf("failed to load config file: %v", err)
	}

	config := Config{}
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadConfig() Config {
	config, err := readConfig(".env")
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	return *config
}
