package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	PostgresDatabase struct {
		URL     string `yaml:"url"`
		DB      string `yaml:"db"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"pgDatabase"`
}

// NewConfig ...
func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {

		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}
