package os

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"syscall"
	"time"

	"dev.azure.com/humana/pie/lib/process/builder"
	"dev.azure.com/humana/pie/lib/process/static"
	"dev.azure.com/humana/pie/lib/process/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// Processor defiation
type Processor struct{}

// P defination
func (p *Processor) P(b builder.Builder) types.Process {
	return &Process{
		Builder: *b.Copy(),
		cmdC:    make(chan *exec.Cmd, 1),
		inPipeC: make(chan io.WriteCloser, 1),
		doneC:   make(chan struct{}, 1),
		exitC:   make(chan int, 1),
		stderr:  nil,
		stdout:  nil,
		stdin:   nil,
	}
}

// Process define
type Process struct {
	builder.Builder
	cmdC    chan *exec.Cmd
	stdout  io.Writer
	stderr  io.Writer
	stdin   io.Reader
	inPipeC chan io.WriteCloser
	doneC   chan struct{}
	exitC   chan int
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
	terminated := make(chan struct{})

	cmd := <-p.cmdC
	inPipe := <-p.inPipeC

	go func() { <-p.doneC; terminated <- struct{}{} }()

	// For some of process, Close stdin will let it terminate
	// 	A caller need only call Close to force the pipe to close sooner.
	// For example, if the command being run will not exit until standard input
	// is closed, the caller must close the pipe.
	inPipe.Close()

	select {
	case <-terminated:
		log.Infof("Process %s With Pid %d terminated after close stdin", cmd, cmd.Process.Pid)
		return
	case <-time.After(1 * time.Second):
	}

	err := cmd.Process.Signal(syscall.SIGINT)
	if err != nil {
		log.Infof("Process %s With Pid %d couldn't terminated after SIGINT", cmd, cmd.Process.Pid)
	}

	select {
	case <-terminated:
		log.Infof("Process %s With Pid %d terminated after SIGINT", cmd, cmd.Process.Pid)
		return
	case <-time.After(2 * time.Second):
	}

	err = cmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		log.Infof("Process %s With Pid %d couldn't terminated after SIGTERM", cmd, cmd.Process.Pid)
	}

	select {
	case <-terminated:
		log.Infof("Process %s With Pid %d terminated after SIGTERM", cmd, cmd.Process.Pid)
		return
	case <-time.After(3 * time.Second):
	}

	err = cmd.Process.Kill()
	if err != nil {
		log.Infof("Process %s With Pid %d couldn't terminated after SIGKILL", cmd, cmd.Process.Pid)
	}

	select {
	case <-terminated:
		log.Infof("Process %s With Pid %d terminated after SIGKILL", cmd, cmd.Process.Pid)
		return
	case <-time.After(4 * time.Second):
	}

	log.Infof("Process %s With Pid %v couldn't terminated!", cmd, cmd.Process.Pid)
}

// Exec start a new Process using the attributes of this process builder
// This method checks that the command is a valid operating system command.
// Which commands are valid is system-dependent,
// but at the very least the command must be a non-empty list of non-null strings
// Among the many things that can go wrong are:
// - The Operating system program file was not found
// - Access the the program file was denied
// - The working directory does not exist.
// - etc
// Each Process can only be executed once
// The behavior of Process that is executed multiple times is undefined
func (p Process) Exec(ctx context.Context) (err error) {
	outWriter := p.stdout
	errWriter := p.stderr

	// check binary from path
	_, err = exec.LookPath(p.Builder.Name)

	if err != nil {
		p.doneC <- struct{}{}

		return errors.Wrapf(err, "couldn't find command %s in path or file is not executable", p.Builder.Name)
	}

	// setup command parameters
	cmd := exec.CommandContext(ctx, p.Builder.Name, p.Builder.Args...)
	p.cmdC <- cmd

	cmd.Dir = p.Builder.WorkDir

	cmd.Env = static.EnvMapToStringSlice(p.Builder.Env)

	inPipe, err := cmd.StdinPipe()
	p.inPipeC <- inPipe

	if err != nil {
		p.doneC <- struct{}{}

		return errors.Wrapf(err, "couldn't get StdinPipe of command %s", cmd)
	}

	outPipe, err := cmd.StdoutPipe()

	if err != nil {
		p.doneC <- struct{}{}
		return errors.Wrapf(err, "couldn't get StdoutPipe of command %s", cmd)
	}

	errPipe, err := cmd.StderrPipe()

	if err != nil {
		p.doneC <- struct{}{}
		return errors.Wrapf(err, "couldn't get StderrPipe of command %s", cmd)
	}

	err = cmd.Start()

	if err != nil {
		p.doneC <- struct{}{}

		return errors.Wrapf(err, "couldn't start command %s process", cmd)
	}

	// Concurrent replication stream
	all, ioctx := errgroup.WithContext(ctx)

	if p.stdin != nil && inPipe != nil {
		all.Go(func() error {
			_, err := static.CancellableCopy(ioctx, inPipe, p.stdin)

			if err != nil {
				return errors.Wrap(err, "copy stdin err")
			}

			return nil
		})
	}

	if outWriter != nil && outPipe != nil {
		all.Go(func() error {
			_, err := static.CancellableCopy(ioctx, outWriter, outPipe)

			if err != nil {
				return errors.Wrap(err, "copy stdout err")
			}

			return nil
		})
	}

	if errWriter != nil && errPipe != nil {
		all.Go(func() error {
			_, err := static.CancellableCopy(ioctx, errWriter, errPipe)
			if err != nil {
				return errors.Wrap(err, "copy stderr err")
			}

			return nil
		})
	}

	err = all.Wait()

	if err != nil {
		p.doneC <- struct{}{}

		return errors.Wrap(err, fmt.Sprintf(
			"command %s running with err, io copy err",
			cmd,
		))
	}

	err = cmd.Wait()

	if err != nil || cmd.ProcessState.ExitCode() != 0 {
		p.doneC <- struct{}{}
		p.exitC <- cmd.ProcessState.ExitCode()
		return errors.Wrapf(err, "command %s finish with err, exit %d", cmd, cmd.ProcessState.ExitCode())
	}

	p.doneC <- struct{}{}
	p.exitC <- cmd.ProcessState.ExitCode()
	return nil
}
