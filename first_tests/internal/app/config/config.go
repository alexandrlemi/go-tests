package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
	JWTKey  string `yaml:"jwt_key"`
}

func LoadConfig(filepath string) (*Config, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("Error reading config file: %s", err)
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Printf("Error parsing config file: %s", err)
		return nil, err
	}

	return &config, nil
}
