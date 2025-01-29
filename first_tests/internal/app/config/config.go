package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
	Auth    struct {
		JWTKey string `yaml:"jwt_key"`
		Pepper string `yaml:"pepper"`
	} `yaml:"auth"`
}

func LoadConfig(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("Error reading config file: %s", err)
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Printf("Error parsing config file: %s", err)
		return nil, err
	}

	return &cfg, nil
}
