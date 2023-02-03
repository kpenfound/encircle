package main

import (
	"context"
	"fmt"
	"io"
	"strings"

	"dagger.io/dagger"
)

type Executor struct {
	Ctx    context.Context
	Client *dagger.Client
	Logger io.Writer
}

func NewExecutor(ctx context.Context, logger io.Writer) (*Executor, error) {
	c, err := dagger.Connect(ctx, dagger.WithLogOutput((logger)))
	return &Executor{
		Ctx:    ctx,
		Logger: logger,
		Client: c,
	}, err
}

func (e *Executor) ExecuteJob(name string, job *Job) error {
	src := e.Client.Host().Directory(".")

	runner := e.Client.Container().
		Pipeline(name).
		From(job.Docker[0].Image).
		WithMountedDirectory("/src", src).
		WithWorkdir("/src")

	for _, s := range job.Steps {
		runner = runner.
			Pipeline(s.Run.Name).
			WithExec(strings.Split(s.Run.Command, " "))
	}
	_, err := runner.ExitCode(e.Ctx)
	return err
}

func (e *Executor) ExecuteWorkflow(name string, workflow *Workflow, jobs map[string]*Job) error {
	e.Logger.Write([]byte(fmt.Sprintf("Running workflow %s\n", name)))

	for _, jobName := range workflow.Jobs {
		err := e.ExecuteJob(jobName, jobs[jobName])
		if err != nil {
			return err
		}
	}
	return nil
}
