package mem

import (
	"context"
	"io"

	"dev.azure.com/humana/pie/lib/process/builder"
	"dev.azure.com/humana/pie/lib/process/static"
	"dev.azure.com/humana/pie/lib/process/types"
	log "github.com/sirupsen/logrus"
)

// Processor defiation
type Processor struct{}

// P defination
func (p *Processor) P(b builder.Builder) types.Process {
	return &Process{
		Builder:       *b.Copy(),
		TerminateFunc: func() { log.Infof("Terminate\n") },
		ExecFunc:      func(_ context.Context) error { log.Infof("Exec\n"); return nil },
		stderr:        nil,
		stdout:        nil,
		stdin:         nil,
	}
}

// Process define
type Process struct {
	builder.Builder
	TerminateFunc func()
	ExecFunc      func(context.Context) error
	stdout        io.Writer
	stderr        io.Writer
	stdin         io.Reader
}

// WithStdout set this ProcessBuilder's stdout stream
func (p *Process) WithStdout(stdout io.Writer) types.Process {
	p.stdout = stdout
	return p
}

// WithStderr set this ProcessBuilder's stderr stream
func (p *Process) WithStderr(stderr io.Writer) types.Process {
	p.stderr = stderr
	return p
}

// WithStdin set this ProcessBuilder's stdin stream
func (p *Process) WithStdin(stdin io.Reader) types.Process {
	p.stdin = stdin
	return p
}

// Name is the path of the command to run.
func (p *Process) Name() string {
	return p.Builder.Name
}

// Args holds command line arguments.
func (p *Process) Args() []string {
	args := make([]string, len(p.Builder.Args))
	copy(args, p.Builder.Args)
	return args
}

// WorkDir specifies the working directory of the command.
func (p *Process) WorkDir() string {
	return p.Builder.WorkDir
}

// Env specifies the environment of the process.
func (p *Process) Env() map[string]string {
	return static.CopyStringMap(p.Builder.Env)
}

// Stdout and Stderr specify the process's standard output and error.
func (p *Process) Stdout() io.Writer {
	return p.stdout
}

// Stderr define
func (p *Process) Stderr() io.Writer {
	return p.stderr
}

// Stdin specifies the process's standard input.
func (p *Process) Stdin() io.Reader {
	return p.stdin
}

// Terminate a process define
func (p *Process) Terminate() {
	p.TerminateFunc()
}

// Exec define
func (p Process) Exec(ctx context.Context) (err error) {
	return p.ExecFunc(ctx)
}
