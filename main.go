package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//go:embed defaultConfig.json
var defaultConfig []byte

func main() {
	  flags := parseFlags()

		// Attempt to read the config file
		config, err := os.ReadFile(flags.ConfigPath)
		if err != nil {
			// If using default path, create the default config there.
        if flags.ConfigPath == "~/.config/palettro/config.json" {
            log.Println("No file found at default path, creating default config.")
            
            homeDir, err := os.UserHomeDir()
            if err != nil {
                log.Fatalf("Unable to get home directory: %v\n", err)
            }
            configPath := filepath.Join(homeDir, ".config", "palettro", "config.json")
            
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

		fmt.Println(string(config))
}