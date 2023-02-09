package main

import (
	"context"
	"fmt"
	"os"
)

// Input options:
// `./encircle`: run all workflows
// `./encircle workflow foo`: run workflow foo
// `./encircle job bar`: run job bar

func main() {
	// Parse input
	mode := "unknown"
	target := ""
	if len(os.Args) == 1 {
		mode = "all"
	} else if len(os.Args) == 3 {
		mode = os.Args[1]
		target = os.Args[2]
	}
	// Load config
	configPath := "./.circleci/config.yml"
	fmt.Printf("loading config at %s\n", configPath)
	cfg, err := readConfig(configPath)
	if err != nil {
		panic(err)
	}

	// Create executor
	executor, err := NewExecutor(context.Background(), os.Stdout)
	if err != nil {
		panic(err)
	}

	// Run target
	switch mode {
	case "all":
		for k, w := range cfg.Workflows {
			err = executor.ExecuteWorkflow(k, w, cfg.Jobs)

		}
	case "workflow":
		err = executor.ExecuteWorkflow(target, cfg.Workflows[target], cfg.Jobs)
		if err != nil {
			panic(err)
		}
	case "job":
		err := executor.ExecuteJob(target, cfg.Jobs[target])
		if err != nil {
			panic(err)
		}
	default:
		fmt.Printf("error: unknown input %+v", os.Args[1:])
	}

	if err != nil {
		panic(err)
	}
}
