// A generated module for DaggerExperiment functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/dagger-experiment/internal/dagger"
)

type DaggerExperiment struct{}

// Returns a container that echoes whatever string argument is provided
func (m *DaggerExperiment) ContainerEcho(ctx context.Context, stringArg string) (string, error) {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg}).Stdout(ctx)
}

func (m *DaggerExperiment) Publish(ctx context.Context) (string, error) {
	return dag.Container().From("alpine:latest").WithNewFile("/hi.txt", "Hello Dagger!\n").WithEntrypoint([]string{"cat", "/hi.txt"}).Publish(ctx, "ttl.sh/hello")
}

func (m *DaggerExperiment) Basic(ctx context.Context) *dagger.Container {
	// aptCache := dag.CacheVolume("apt-cache")
	return dag.Container().From("debian:latest"). /*.WithMountedCache("/var/cache/apt/archives", aptCache)*/ WithExec([]string{"apt-get", "update"}).WithExec([]string{"apt-get", "install", "--yes", "maven", "mariadb-server"})
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *DaggerExperiment) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}
