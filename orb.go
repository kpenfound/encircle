package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	orbYaml, err := queryOrbDetails(value.Value)
	if err != nil {
		return err
	}
	var orb *OrbConfig
	err = yaml.Unmarshal([]byte(orbYaml), &orb)
	if err != nil {
		return err
	}
	o.Orb = orb
	o.Name = value.Value
	return nil
}

func queryOrbDetails(versionref string) (string, error) {
	name := strings.Split(versionref, "@")[0]
	payload := fmt.Sprintf(`
		{"operationName":"OrbDetailsQuery",
		"variables":{"name":"%s","orbVersionRef":"%s"},
		"query":"query OrbDetailsQuery($name: String, $orbVersionRef: String) {\n
				orbVersion(orbVersionRef: $orbVersionRef) {\n
					source\n
				}\n
			}\n
		"}
	`, name, versionref)
	endpoint := "https://circleci.com/graphql-unstable"

	body := bytes.NewReader([]byte(payload))
	resp, err := http.Post(endpoint, "application/json", body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	source := struct {
		Data struct {
			OrbVersion struct {
				Source string `json:"source"`
			} `json:"orbVersion"`
		} `json:"data"`
	}{}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&source)
	if err != nil {
		return "", err
	}
	return source.Data.OrbVersion.Source, nil
}
