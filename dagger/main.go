// A generated module for Daggertest functions
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
	"debug/elf"
	"fmt"
	"os"
	"path/filepath"
)

type Daggertest struct{}

const exampleMain = `package main

import "fmt"

func main() {
	fmt.Println("hello from arm64..")
}
`

// Cross-compiles a simple Go example project for linux/arm64 and verifies the
// resulting binary is an AArch64 ELF executable.
//
// +check
func (m *Daggertest) Test(ctx context.Context) error {
	binary := dag.Container().
		From("golang:1.26-alpine").
		WithWorkdir("/src").
		WithNewFile("go.mod", "module example\n\ngo 1.26\n").
		WithNewFile("main.go", exampleMain).
		WithEnvVariable("CGO_ENABLED", "0").
		WithEnvVariable("GOOS", "linux").
		WithEnvVariable("GOARCH", "arm64").
		WithExec([]string{"go", "build", "-o", "/out/example", "."}).
		File("/out/example")

	dir, err := os.MkdirTemp("", "example")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	path := filepath.Join(dir, "example")
	if _, err := binary.Export(ctx, path); err != nil {
		return err
	}

	f, err := elf.Open(path)
	if err != nil {
		return fmt.Errorf("parse binary as ELF: %w", err)
	}
	defer f.Close()

	if f.Machine != elf.EM_AARCH64 {
		return fmt.Errorf("expected %s binary, got %s", elf.EM_AARCH64, f.Machine)
	}
	return nil
}
