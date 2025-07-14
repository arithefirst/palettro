package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath) // Attempt to get file information
	if err != nil {
		// If an error occurred, check if it indicates the file doesn't exist
		return !errors.Is(err, os.ErrNotExist)
	}
	// If no error, the file exists
	return true
}

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

	color, colorExists := config.Colors[flags.Color]

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

		for _, f := range dir {
			filePath := filepath.Join(path, f.Name())
			file, err := os.ReadFile(filePath)
			if err != nil {
				log.Fatalf("[ENOENT]: Unable to read file at \"%s\"\n", filePath)
			}

			var fileStr string
			fileStr = strings.ReplaceAll(string(file), "((PALETTRO.HEX))", color.Hex)
			fileStr = strings.ReplaceAll(fileStr, "((PALETTRO.HSL))", color.HSL)
			fileStr = strings.ReplaceAll(fileStr, "((PALETTRO.RGB))", color.RGB)
			fileStr = strings.ReplaceAll(fileStr, "((PALETTRO.RGBA))", color.RGBA)
			fileStr = strings.ReplaceAll(fileStr, "((PALETTRO.HEXTRANS))", color.HexTrans)

			newFilePath := expandPath(filepath.Join(v.Path, f.Name()))

			if fileExists(newFilePath) && !flags.Autoconfirm {
				fmt.Printf("Warning: File '%s' already exists! Continuing will overwrite it. Continue? [y/N]: ", filePath)

				var response string
				fmt.Scan(&response)

				response = strings.ToLower(strings.TrimSpace(response))
				if !(response == "y" || response == "yes") {
					os.Exit(1)
				}
			}

			err = os.WriteFile(newFilePath, []byte(fileStr), 0644)
			if err != nil {
				log.Fatalf("Failed to write file '%s': %v", newFilePath, err)
			}
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
	}
}
