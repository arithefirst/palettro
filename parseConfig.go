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

// Holds info for a defined "color"
type Color struct {
	Hex string `json:"hex"`
	RGB string `json:"rgb"`
	HSL string `json:"hsl"`
}

// Holds info about a config that needs to be updated
type ServiceConfig struct {
	Path   string `json:"path"`
	Restart string `json:"restart,omitempty"`
	Name string `json:"name"`
}

type Config struct {
	Colors      map[string]Color `json:"colors"`
	ConfigFiles []ServiceConfig     `json:"configs"`
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

	// Validate that all color names are lowercase
	for key := range parsedConfig.Colors {
		if key != strings.ToLower(key) {
			log.Fatalf("Config error: Color name '%s' must be lowercase", key)
		}
	}

	for _,v := range parsedConfig.ConfigFiles {
		os.MkdirAll(expandPath("~/.config/palettro/" + strings.ToLower(v.Name)), 0755)
	}

	return parsedConfig
}