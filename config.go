package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Client_id string `yaml:"CLIENT_ID"`
	Api_token string `yaml:"API_TOKEN"`
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
