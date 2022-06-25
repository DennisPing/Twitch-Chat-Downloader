package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ClientId string `yaml:"client_id"`
	ApiToken string `yaml:"api_token"`
	Format   struct {
		Comments struct {
			Badges struct {
				Admin       string `yaml:"admin"`
				Bits        string `yaml:"bits"`
				Broadcaster string `yaml:"broadcaster"`
				GlobalMod   string `yaml:"global_mod"`
				Moderator   string `yaml:"moderator"`
				Premium     string `yaml:"premium"`
				Staff       string `yaml:"staff"`
				Subscriber  string `yaml:"subscriber"`
				Turbo       string `yaml:"turbo"`
			} `yaml:"badges"`
			Format    string `yaml:"format"`
			Timestamp struct {
				Relative string `yaml:"relative"`
			} `yaml:"timestamp"`
		} `yaml:"comments"`
		Filename struct {
			Format    string `yaml:"format"`
			Timestamp struct {
				Absolute string `yaml:"absolute"`
			} `yaml:"timestamp"`
		} `yaml:"filename"`
	} `yaml:"format"`
}

// Parse the config.yml file and return a Config struct.
func ParseConfig(file string) (*Config, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	content = []byte(os.ExpandEnv(string(content)))
	config := &Config{}
	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Check for the config.yml file and create one if not exists.
func LoadConfig() {
	if config_dir, err := os.UserConfigDir(); err == nil {
		// On Windows: %APPDATA%/tcd-go/config.yml
		// On Linux: $XDG_CONFIG_HOME/.config/tcd-go/config.yml
		// On Mac: $HOME/Library/Application Support/tcd-go/config.yml
		app_dir := filepath.Join(config_dir, "tcd-go")
		config_path := filepath.Join(app_dir, "config.yml")
		if _, err := os.Stat(config_path); err != nil {
			fmt.Printf("Creating config at: %s\n", config_path)
			err := createConfig(app_dir)
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

// Create the config.yaml config file.
func createConfig(config_dir string) error {
	err := os.MkdirAll(config_dir, 0755)
	if err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(config_dir, "config.yml"))
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
