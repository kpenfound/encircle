package main

import (
	"io/ioutil"

	"dagger.io/dagger"
	"gopkg.in/yaml.v3"
)

type Orb struct {
	Name string
	Orb  *OrbConfig
}

type OrbConfig struct {
	Version     string `yaml:"version"`
	Description string `yaml:"description"`
	//	Jobs      map[string]*OrbJob         `yaml:"jobs"` // TODO
	Commands map[string]*OrbCommand `yaml:"commands"`
	// Executors map[string]*OrbExecutor `yaml:"executors"` // TODO
}

type OrbCommand struct {
	Steps      []*Step                  `yaml:"steps"`
	Parameters map[string]*OrbParameter `yaml:"parameters"`
}

type OrbParameter struct {
	DefaultValue string   `yaml:"default"`
	Description  string   `yaml:"description"`
	ParamType    string   `yaml:"type"`
	Enum         []string `yaml:"enum"`
}

// type OrbJob struct{} // TODO
// type OrbExecutor struct{}

func (oc *OrbCommand) ToDagger(c *dagger.Container, params map[string]string) *dagger.Container {
	// TODO: handle params?
	for _, s := range oc.Steps {
		c = s.ToDagger(c, params)
	}
	return c
}

func (oc *OrbCommand) GetDefaultParameters() map[string]string {
	defaults := map[string]string{}
	for k, v := range oc.Parameters {
		defaults[k] = v.DefaultValue
	}
	return defaults
}

func (o *Orb) UnmarshalYAML(value *yaml.Node) error {
	if value.Tag == "!!str" {
		// TODO: get orb yaml and parse it into Jobs, Commands, Executors
		orbYaml, err := ioutil.ReadFile(".orb_test.yml") // TODO: fetch this remotely
		if err != nil {
			return err
		}
		var orb *OrbConfig
		err = yaml.Unmarshal(orbYaml, &orb)
		if err != nil {
			return err
		}
		o.Orb = orb
		o.Name = value.Value
	}
	return nil
}
