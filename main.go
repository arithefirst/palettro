package main

import (
	"fmt"
	"log"
)

func main() {
	  flags := parseFlags()
		config := parseConfig(flags)

		color, colorExists := config.Colors[flags.Color]

		if flags.Color == "N/A" {
			log.Fatalln("The \"-color\" flag must be set.")
		} else if !colorExists {
			log.Fatalf("The Color \"%s\" does not exist in your config (%s)", flags.Color, flags.ConfigPath)
		}

		fmt.Println(color)
}