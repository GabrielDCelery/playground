package main

import (
	"context"
	"dagger/dagger-experiment/internal/dagger"
)

type DaggerExperiment struct{}

func (m *DaggerExperiment) BuildApp(ctx context.Context, src *dagger.Directory) *dagger.File {
	return dag.Container().
		From("golang:1.24-alpine").
		WithDirectory("/app", src).
		WithWorkdir("/app").
		WithEnvVariable("GOARCH", "amd64").
		WithEnvVariable("GOOS", "linux").
		WithEnvVariable("CGO_ENABLED", "0").
		WithExec([]string{"go", "build", "-o", "./goserver"}).
		File("/app/goserver")
}
