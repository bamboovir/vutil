package builder

import (
	"dev.azure.com/humana/pie/lib/process/static"
)

// Builder manage a collection of process attributes
type Builder struct {
	Name    string
	Args    []string
	WorkDir string
	Env     map[string]string
}

// Copy Builder define
func (b Builder) Copy() *Builder {
	copyB := &Builder{}
	copyB.Name = b.Name
	copyB.WorkDir = b.WorkDir

	copyB.Args = make([]string, len(b.Args))
	copy(copyB.Args, b.Args)
	copyB.Env = static.CopyStringMap(b.Env)

	return copyB
}

// New is a empty constructor
// construct a Builder
func New(name string) *Builder {
	return &Builder{
		Name:    name,
		Args:    []string{},
		WorkDir: "/",
		Env:     map[string]string{},
	}
}

// WithName define
func (b *Builder) WithName(name string) *Builder {
	b.Name = name
	return b
}

// WithArgs constructs a Builder with a string list
// containing the same strings as the command array.
// It's not checked whether command corresponds to a valid operating
// system command
func (b *Builder) WithArgs(args []string) *Builder {
	b.Args = make([]string, len(args))
	copy(b.Args, args)
	return b
}

// AppendArgs define
func (b *Builder) AppendArgs(args []string) *Builder {
	for _, arg := range args {
		b.Args = append(b.Args, arg)
	}
	return b
}

// WithWorkDir set this Processbuilder's working directory.
// Subprocesses subsequently started by Exec method will use this
// as their working dircctory.
func (b *Builder) WithWorkDir(workDir string) *Builder {
	b.WorkDir = workDir
	return b
}

// WithEnv define
func (b *Builder) WithEnv(env map[string]string) *Builder {
	b.Env = static.CopyStringMap(env)
	return b
}

// WithEnvAppend define
func (b *Builder) WithEnvAppend(env map[string]string) *Builder {
	b.Env = static.MergeEnvMap(b.Env, env)
	return b
}
