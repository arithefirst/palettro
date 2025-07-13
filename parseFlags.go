package main

import "flag"

type Flags struct {
	ConfigPath string
	Color string
}

func parseFlags() Flags {
	var returnFlags Flags;
	
	// Get flags
	configPath := flag.String("config", "~/.config/palettro/config.json", "Sets the path to the config file")
	color := flag.String("color", "N/A", "New accent color to use for your RICE")

	flag.Parse()
	
	// Populate struct
	returnFlags.ConfigPath = *configPath
	returnFlags.Color = *color

	return returnFlags
}