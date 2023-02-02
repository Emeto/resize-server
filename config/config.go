package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App     `json:"app"`
		HTTP    `json:"http"`
		Log     `json:"logger"`
		Presets []Preset `json:"presets"`
	}

	App struct {
		Name    string `env-required:"true" json:"name" env:"APP_NAME"`
		Version string `env-required:"true" json:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" json:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" json:"logLevel" env:"LOG_LEVEL"`
	}

	Preset struct {
		Name   string `json:"name"`
		Width  uint   `json:"width"`
		Height uint   `json:"height"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.json", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
