package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	  flags := parseFlags()
		config := parseConfig(flags)

		if flags.ShowColors {
			for i := range config.Colors {
				fmt.Println(i)
			}
			os.Exit(0)
		}

		if flags.ShowConfigs {
			fmt.Println("All Configs:")
			for _,v := range config.ConfigFiles {
				fmt.Println("----------------")
				fmt.Printf("Name: %s\n", v.Name)
				fmt.Printf("Location (palettro): ~/.config/palettro/%s\n", strings.ToLower(v.Name))
				fmt.Printf("Location: %s\n", v.Path)
				fmt.Printf("Restart service on color change: %t\n", v.Restart != "")
				if v.Restart != "" {
					fmt.Printf("Service to kill on color change: %v\n", v.Restart)
				}
			}
			os.Exit(0)
		}

		_, colorExists := config.Colors[flags.Color]

		if flags.Color == "N/A" {
			log.Fatalln("The \"-color\" flag must be set.")
		} else if !colorExists {
			log.Fatalf("The Color \"%s\" does not exist in your config (%s)", flags.Color, flags.ConfigPath)
		}

		// To-Do: 
		// - Read config files for each config setting in the config.json (EX: if config name is waybar, 
		// it gets mapped to ~/.config/palettro/waybar/)
		// - Add main functionality of writing configs to the correct folders in the system
		// - Add functionality of service restarting for services that don't auto-update configs
}