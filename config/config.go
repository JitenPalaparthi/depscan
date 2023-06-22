package config

import (
	_ "embed"
	"encoding/json"
)

//go:embed config.json
var configFile string

type Config struct {
	IgnoreDirs  []string     `json:"ignoreDirs"`
	IgnoreFiles []string     `json:"ignoreFiles"`
	DepManagers []DepManager `json:"depManagers"`
}

func New() (*Config, error) {
	config := new(Config)
	err := json.Unmarshal([]byte(configFile), config)
	return config, err
}

func (c *Config) GetExtensions() []string {
	exts := make([]string, len(c.DepManagers))
	for i, v := range c.DepManagers {
		exts[i] = v.FileExt
	}
	return exts
}

func (c *Config) GetDepFiles() []string {
	fileNames := make([]string, len(c.DepManagers))
	for i, v := range c.DepManagers {
		fileNames[i] = v.FileName
	}
	return fileNames
}

func (c *Config) GetDepManagerByExt(ext string) *DepManager {
	for i, v := range c.DepManagers {
		if v.FileExt == ext {
			return &c.DepManagers[i]
		}
	}
	return nil
}

func (c *Config) GetDepManagerByFileName(fileName string) *DepManager {
	for i, v := range c.DepManagers {
		if v.FileName == fileName {
			return &c.DepManagers[i]
		}
	}
	return nil
}
