package main

import (
	"path/filepath"

	"gopkg.in/ini.v1"
)

func configure(c, v string) error {
	cfgDir := filepath.Join(REPO_DIR, CONFIG_FILE)
	cfg, err := ini.Load(cfgDir)
	if err != nil {
		return err
	}

	switch c {
	case "user.email":
		cfg.Section("user").Key("email").SetValue(v)
	case "user.name":
		cfg.Section("user").Key("name").SetValue(v)
	default:
		print("Section not found")
		return nil
	}

	cfg.SaveTo(cfgDir)

	return nil
}
