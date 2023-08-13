package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	API API `yaml:"api"`
}

type API struct {
	Port           string        `yaml:"port"`
	MaxHeaderBytes int           `yaml:"maxHeaderBytes"`
	Timeout        time.Duration `yaml:"timeout"`
}

func NewConfig() (*Config, error) {
	var cfg *Config

	//unmarshal config into "config" struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error while parsing configs: %w", err)
	}

	return cfg, nil
}
