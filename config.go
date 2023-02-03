package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version   string               `yaml:"version"`
	Jobs      map[string]*Job      `yaml:"jobs"`
	Workflows map[string]*Workflow `yaml:"workflows"`
}

type Job struct {
	Docker []*Docker `yaml:"docker"`
	Steps  []*Step   `yaml:"steps"`
}

type Workflow struct {
	Jobs []string `yaml:"jobs"`
}

type Docker struct {
	Image string `yaml:"image"`
}

type Step struct {
	NamedTask string `yaml:",omitempty"`
	Run       struct {
		Name    string `yaml:"name"`
		Command string `yaml:"command"`
	} `yaml:"run,omitempty"`
}

func readConfig(path string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var configParsed *Config
	err = yaml.Unmarshal(configBytes, &configParsed)
	return configParsed, err
}
