package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

type User struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

type Config struct {
	User User
}

func configure(c, v string) error {
	cfg, err := readFromConfig()
	if err != nil {
		return err
	}

	switch c {
	case "user.email":
		cfg.User.Email = v
	case "user.name":
		cfg.User.Name = v
	default:
		print("Section not found")
		return nil
	}

	err = writeToConfig(cfg)
	if err != nil {
		return err
	}

	return nil
}

func writeToConfig(cfg *Config) error {
	data, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(cfgDir(), data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func readFromConfig() (*Config, error) {
	var config Config

	f, err := os.Open(cfgDir())
	if err != nil {
		return nil, err
	}

	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func getUserData() (*User, error) {
	config, err := readFromConfig()
	if err != nil {
		return nil, err
	}

	return &config.User, nil
}

func cfgDir() string {
	return filepath.Join(REPO_DIR, CONFIG_FILE)
}
