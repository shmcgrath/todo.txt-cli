package config

import (
	"encoding/json"
	"log"
	"os"
)

func Load(configPath string) (*Config, error) {
	byt, err := os.ReadFile(os.ExpandEnv(configPath))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(byt, &cfg); err != nil {
		log.Println(err)
		return nil, err
	}

	cfg.CfgPath = configPath

	return &cfg, nil
}
