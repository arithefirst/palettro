package main

import "flag"

type Flags struct {
	ConfigPath string
}

func parseFlags() Flags {
	var returnFlags Flags;
	
	// Get flags
	configPath := flag.String("c", "~/.config/palettro/config.json", "Sets the path to the config file")
	flag.Parse()
	
	// Populate struct
	returnFlags.ConfigPath = *configPath

	return returnFlags
}