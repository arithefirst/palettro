package main

import "flag"

type Flags struct {
	ConfigPath  string
	Color       string
	ShowColors  bool
	ShowConfigs bool
	Autoconfirm bool
}

func parseFlags() Flags {
	var returnFlags Flags

	// Get flags
	configPath := flag.String("config", "~/.config/palettro/config.json", "Sets the path to the config file")
	color := flag.String("color", "N/A", "New accent color to use for your RICE")
	showColors := flag.Bool("showcolors", false, "Show the names of all colors registered in the config")
	showConfigs := flag.Bool("showconfigs", false, "Show all the names and details of registered configs")
	autoconfirm := flag.Bool("autoconfirm", false, "Skip warning about file overwrites")

	flag.Parse()

	// Populate struct
	returnFlags.ConfigPath = *configPath
	returnFlags.Color = *color
	returnFlags.ShowColors = *showColors
	returnFlags.ShowConfigs = *showConfigs
	returnFlags.Autoconfirm = *autoconfirm

	return returnFlags
}
