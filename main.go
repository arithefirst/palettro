package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
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
		for _, v := range config.ConfigFiles {
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

	for _, v := range config.ConfigFiles {
		path := expandPath("~/.config/palettro/" + strings.ToLower(v.Name))
		dir, err := os.ReadDir(path)
		if err != nil {
			log.Fatalf("[ENOENT]: Unable to read directory at %v\n", path)
		}

		if v.Restart != "" {
			// Kill the process by name
			killCmd := exec.Command("pkill", "-f", v.Restart)
			if err := killCmd.Run(); err != nil {
				log.Printf("Warning: Failed to kill process %s: %v", v.Restart, err)
			}

			// Restart the process detached from this program
			restartCmd := exec.Command("nohup", v.Restart)
			restartCmd.SysProcAttr = &syscall.SysProcAttr{
				Setpgid: true,
			}
			if err := restartCmd.Start(); err != nil {
				log.Printf("Warning: Failed to restart process %s: %v", v.Restart, err)
			}
		}

		fmt.Println(dir)
	}

	// To-Do:
	// - Add main functionality of writing configs to the correct folders in the system
}
