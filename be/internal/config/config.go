package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	PortHTTP int `env:"port_http"`
}

func MustLoadConfig(relativePath string) *Config {
	_, err := os.Stat(relativePath)
	if os.IsNotExist(err) {
		panic("Config file is not exist: " + relativePath)
	}

	var cfg Config

	err = cleanenv.ReadConfig(relativePath, cfg)
	if err != nil {
		panic("Config file could not be read: " + relativePath)
	}

	return &cfg
}
