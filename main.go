package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	// Load config
	configPath := "./.circleci/config.yml"
	fmt.Printf("Loading config at %s\n", configPath)
	cfg, err := readConfig(configPath)
	if err != nil {
		panic(err)
	}

	// Create executor
	executor, err := NewExecutor(context.Background(), os.Stdout)
	if err != nil {
		panic(err)
	}

	// Run workflow
	workflow := "build_test"
	err = executor.ExecuteWorkflow(workflow, cfg.Workflows[workflow], cfg.Jobs)
	if err != nil {
		panic(err)
	}
}
