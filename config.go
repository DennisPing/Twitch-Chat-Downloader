package main

import (
	"io/ioutil"
	"os"

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

func ParseConfig(file string) (*Config, error) {
	content, err := ioutil.ReadFile(file)
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
