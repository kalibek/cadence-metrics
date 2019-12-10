package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Config struct {
	Cadence struct {
		Server   string
		Port     string
		Domain   string
		TaskList []string
	}
}

func LoadConfig(configPath string) (*Config, error) {
	c := &Config{}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
