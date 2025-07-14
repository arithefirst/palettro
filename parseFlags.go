package main

import "flag"

type Flags struct {
	ConfigPath string
	Color string
	ShowColors bool
}

func parseFlags() Flags {
	var returnFlags Flags;
	
	// Get flags
	configPath := flag.String("config", "~/.config/palettro/config.json", "Sets the path to the config file")
	color := flag.String("color", "N/A", "New accent color to use for your RICE")
	showColors := flag.Bool("showcolors", false, "Show the names of all colors registered in the config")

	flag.Parse()
	
	// Populate struct
	returnFlags.ConfigPath = *configPath
	returnFlags.Color = *color
	returnFlags.ShowColors = *showColors

	return returnFlags
}