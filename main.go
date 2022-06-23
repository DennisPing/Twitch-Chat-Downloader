package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func cliUsage() {
	fmt.Printf("Usage: %s [-v] URL\n", os.Args[0])
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = cliUsage
	input := parseArgs()
	// This input could either be a complete URL or a video ID
	get_video_req, err := buildGetVideoReq(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(get_video_req)
	// https://www.reddit.com/r/PHP/comments/gtpm5r/what_are_the_benefits_of_using_env_over_envjson/
	loadConfig()
}

// Parse the args and optional flags and return sanitized input
func parseArgs() string {
	flag.BoolVar(&Verbose, "v", false, "verbose output")
	flag.Parse()

	// Validate input args
	if flag.NArg() == 1 {
		return flag.Arg(0)
	} else if flag.NArg() < 1 {
		fmt.Println("Error: missing URL")
		cliUsage()
		os.Exit(1)
	} else {
		fmt.Printf("Error: too many arguments. Expected 1, got %d\n", flag.NArg())
		cliUsage()
		os.Exit(1)
	}
	return ""
}

// Check for the config.yml file and create one if not exists.
func loadConfig() {
	if config_dir, err := os.UserConfigDir(); err == nil {
		// On Windows: %APPDATA%/tcd-go/settings.yml
		// On Linux: $XDG_CONFIG_HOME/.config/tcd-go/settings.yml
		// On Mac: $HOME/Library/Application Support/tcd-go/settings.yml
		config_path := filepath.Join(config_dir, "settings.yml")
		if _, err := os.Stat(config_path); err != nil {
			err := createConfig(config_dir)
			if err != nil {
				fmt.Printf("unable to create config dir: %v\n", err)
				os.Exit(1)
			}
		}
		err := configToEnv(config_path)
		if err != nil {
			fmt.Printf("unable to load config file: %v\n", err)
			os.Exit(1)
		}
	}
}

// Load the config file into env variables.
func configToEnv(config_path string) error {
	f, err := os.Open(config_path)
	if err != nil {
		return err
	}
	defer f.Close()
	// TODO: parse the config file
	return nil
}

// Create the settings.yaml config file.
func createConfig(config_dir string) error {
	err := os.MkdirAll(config_dir, 0755)
	if err != nil {
		return err
	}
	config_path := filepath.Join(config_dir, "settings.yml")
	f, err := os.Create(config_path)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
