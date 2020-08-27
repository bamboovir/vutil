package types

import (
	"context"
	"io"
)

// Process defination
type Process interface {
	// Name is the path of the command to run.
	Name() string
	// Args holds command line arguments.
	Args() []string
	// WorkDir specifies the working directory of the command.
	WorkDir() string
	// Env specifies the environment of the process.
	Env() map[string]string
	// WithStdout and Stderr specify the process's standard output and error.
	WithStdout(io.Writer) Process
	WithStderr(io.Writer) Process
	// WithStdin specifies the process's standard input.
	WithStdin(io.Reader) Process

	Stdout() io.Writer
	Stderr() io.Writer
	// Stdin specifies the process's standard input.
	Stdin() io.Reader
	// Terminate a process
	Terminate()
	// Exec starts the specified command and waits for it to complete.
	Exec(context.Context) error
}
