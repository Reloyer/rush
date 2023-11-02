package config

import (
	"gopkg.in/ini.v1"
)

type Config struct {
	Language     string
	Theme        string
	WindowWidth  int
	WindowHeight int
	LockfilePath string
}

func LoadConfig(configFilePath string) (*Config, error) {
	cfg, err := ini.Load(configFilePath)
	if err != nil {
		return nil, err
	}

	config := &Config{
		Language:     cfg.Section("general").Key("language").String(),
		Theme:        cfg.Section("general").Key("theme").String(),
		LockfilePath: cfg.Section("general").Key("lockfile").String(),
		WindowWidth:  cfg.Section("window").Key("width").MustInt(800),
		WindowHeight: cfg.Section("window").Key("height").MustInt(600),
	}

	if config.Language == "" {
		config.Language = "en"
	}
	if config.Theme == "" {
		config.Theme = "default"
	}

	return config, nil
}
