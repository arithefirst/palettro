package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//go:embed defaultConfig.json
var defaultConfig []byte

type Color struct {
	Hex string `json:"hex"`
	RGB string `json:"rgb"`
	HSL string `json:"hsl"`
}

type Config struct {
	Colors      map[string]Color `json:"colors"`
	ConfigFiles []string         `json:"configFiles"`
}

// expandPath expands a path that starts with "~/" into the full absolute path
// ex: ~/ would become /home/user/
func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Unable to get home directory: %v\n", err)
		}
		return filepath.Join(homeDir, path[2:])
	}
	return path
}

func parseConfig(flags Flags) Config {
	configPath := expandPath(flags.ConfigPath)
	config, err := os.ReadFile(configPath)
	if err != nil {
		// If using default path, create the default config there.
		if flags.ConfigPath == "~/.config/palettro/config.json" {
			log.Println("No file found at default path, creating default config.")

			if err := os.MkdirAll(filepath.Dir(configPath), 0755); err != nil {
				log.Fatalf("Unable to create config directory: %v\n", err)
			}

			if err := os.WriteFile(configPath, defaultConfig, 0644); err != nil {
				log.Fatalf("Unable to write config file: %v\n", err)
			}

			config = defaultConfig
		} else {
			log.Fatalf("[ENOENT]: Unable to read config file at \"%s\"\n", flags.ConfigPath)
		}
	}

	var parsedConfig Config
	if err := json.Unmarshal(config, &parsedConfig); err != nil {
		log.Fatalf("Unable to parse config file: %v\n", err)
	}

	return parsedConfig
}