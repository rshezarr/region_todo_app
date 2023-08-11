package config

import (
	"github.com/spf13/viper"
	"log"
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

func NewConfig() *Config {
	var cfg *Config

	//unmarshal config into "config" struct
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("error while parsing configs: %v", err)
		// return nil
	}

	return cfg
}
