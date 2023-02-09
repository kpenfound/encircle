package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"strings"

	"dagger.io/dagger"
	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
)

type Step struct {
	Name    string `yaml:"name"`
	Run     *Run   `yaml:"run"`
	Command *OrbCommandExecution
	WorkDir string `yaml:"working_directory"`
}

type Run struct {
	Name        string            `yaml:"name"`
	Command     string            `yaml:"command"`
	Environment map[string]string `yaml:"environment"`
}

type OrbCommandExecution struct {
	OrbCommand *OrbCommand
	Parameters map[string]string
}

func (s *Step) ToDagger(c *dagger.Container, params map[string]string) *dagger.Container {
	c = c.Pipeline(s.Name)
	if s.WorkDir != "" { // workdir relative to project root
		c = c.WithWorkdir(filepath.Join("/src", s.WorkDir))
	}
	if s.Run != nil {
		c = s.Run.ToDagger(c, params)
	}
	if s.Command != nil {
		// Get default params
		maps.Copy(params, s.Command.OrbCommand.GetDefaultParameters())
		// Override user params
		maps.Copy(params, s.Command.Parameters)
		c = s.Command.OrbCommand.ToDagger(c, params)
	}

	return c
}

func (r *Run) ToDagger(c *dagger.Container, params map[string]string) *dagger.Container {
	c = c.Pipeline(ReplaceParams(r.Name, params))
	// Set env vars
	for k, v := range r.Environment {
		c = c.WithEnvVariable(k, ReplaceParams(v, params))
	}
	// Exec command
	command := ReplaceParams(r.Command, params)
	command = fmt.Sprintf("#!/bin/bash\n%s", command)
	script := fmt.Sprintf("/%s.sh", getSha(command))
	c = c.WithNewFile(script, dagger.ContainerWithNewFileOpts{
		Permissions: 0777,
		Contents:    command,
	})
	c = c.WithExec([]string{script})
	// Unset env vars
	for k := range r.Environment {
		c = c.WithoutEnvVariable(k)
	}
	return c
}

func (s *Step) UnmarshalYAML(value *yaml.Node) error {
	switch value.Tag {
	case "!!str": // Basic command like checkout
		if value.Value == "checkout" {
			fmt.Println("warning: skipping checkout for local dev")
		} else if strings.Contains(value.Content[0].Value, "/") {
			// Handle orb command with no params
			commandParts := strings.Split(value.Content[0].Value, "/")
			orb := commandParts[0]
			command := commandParts[1]
			s.Command = &OrbCommandExecution{
				OrbCommand: findCommandForOrb(orb, command),
			}
		} else {
			fmt.Printf("warning: unknown step command: %s\n", value.Value)
		}
	case "!!map":
		if len(value.Content) == 0 {
			break
		}
		// Basic run block
		if value.Content[0].Value == "run" {
			// run block
			if value.Content[1].Tag == "!!map" {
				var r *Run
				err := value.Content[1].Decode(&r)
				if err != nil {
					return err
				}
				s.Run = r
			}
			// inline run
			if value.Content[1].Tag == "!!str" {
				r := &Run{}
				r.Command = value.Content[1].Value
				s.Run = r
			}

			// handle orb command with params
		} else if strings.Contains(value.Content[0].Value, "/") {
			commandParts := strings.Split(value.Content[0].Value, "/")
			orb := commandParts[0]
			command := commandParts[1]
			s.Command = &OrbCommandExecution{
				OrbCommand: findCommandForOrb(orb, command),
			}

			// parse params
			if len(value.Content) > 1 {
				params := map[string]string{}
				for i := 0; i < len(value.Content[1].Content); i += 2 {
					k := value.Content[1].Content[i].Value
					v := value.Content[1].Content[i+1].Value
					params[k] = v
				}
				s.Command.Parameters = params
			}
		} else {
			fmt.Printf("Dont know how to handle command %s\n", value.Content[0].Value)
		}
	default:
		fmt.Printf("Unknown yaml Tag %s\n", value.Tag)
	}
	return nil
}

func findCommandForOrb(orb string, command string) *OrbCommand {
	if Glorbs[orb] != nil {
		return Glorbs[orb].Orb.Commands[command]
	}
	fmt.Println("didnt find orb")
	return nil
}

func ReplaceParams(target string, params map[string]string) string {
	if strings.Contains(target, "<< parameters.") || strings.Contains(target, "<<parameters.") {
		for k, v := range params {
			p1 := fmt.Sprintf("<< parameters.%s >>", k)
			p2 := fmt.Sprintf("<<parameters.%s>>", k)
			target = strings.ReplaceAll(target, p1, v)
			target = strings.ReplaceAll(target, p2, v)
		}
	}
	return target
}

func getSha(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
