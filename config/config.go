package config

import (
	_ "embed"
	"encoding/json"
)

//go:embed config.json
var configFile string

type Config struct {
	IgnoreDirs  []string `json:"ignoreDirs"`
	IgnoreFiles []string `json:"ignoreFiles"`
}

func New() (*Config, error) {
	config := new(Config)
	err := json.Unmarshal([]byte(configFile), config)
	return config, err
}
