package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Global orb list for command matching during unmarshalling
var Glorbs = map[string]*Orb{}

type Config struct {
	Version   string               `yaml:"version"`
	Jobs      map[string]*Job      `yaml:"jobs"`
	Workflows map[string]*Workflow `yaml:"workflows"`
	Orbs      map[string]*Orb      `yaml:"orbs"`
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

// Custom parser because orbs have to be evaluated before jobs
func (c *Config) UnmarshalYAML(value *yaml.Node) error {
	nodes := map[string]*yaml.Node{}
	for i := 0; i < len(value.Content); i += 2 {
		k := value.Content[i]
		if k.Tag == "!!str" {
			nodes[value.Content[i].Value] = value.Content[i+1]
		}
	}

	// config.Version
	if nodes["version"] != nil {
		err := nodes["version"].Decode(&c.Version)
		if err != nil {
			return err
		}
	}

	// config.Orbs
	if nodes["orbs"] != nil {
		err := nodes["orbs"].Decode(&c.Orbs)
		if err != nil {
			return err
		}
	}
	Glorbs = c.Orbs

	// config.Jobs
	if nodes["jobs"] != nil {
		err := nodes["jobs"].Decode(&c.Jobs)
		if err != nil {
			return err
		}
	}

	// config.Workflows
	if nodes["workflows"] != nil {
		err := nodes["workflows"].Decode(&c.Workflows)
		if err != nil {
			return err
		}
	}

	return nil
}

func readConfig(path string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// parse yaml
	var configParsed *Config
	err = yaml.Unmarshal(configBytes, &configParsed)

	return configParsed, err
}
