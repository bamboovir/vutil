/*
Package os is mainly used to abstract constant configuration in the life cycle of a single process (
	mock for environment and filesystem is not within the scope of responsibility of this package
)
*/
package os

import (
	"dev.azure.com/humana/pie/lib/os/mem"
	"dev.azure.com/humana/pie/lib/os/os"
)

// OS interface defination
type OS interface {
	Executable() (string, error)
	UserHomeDir() (string, error)
	UserConfigDir() (string, error)
	UserCacheDir() (string, error)
	Getegid() int
	Geteuid() int
	Getgid() int
	Getgroups() ([]int, error)
	Getpagesize() int
	Getpid() int
	Getppid() int
	Getuid() int
	Getwd() (dir string, err error)
}

// NewOS defination
func NewOS() OS {
	return &os.OS{}
}

// NewMem defination
func NewMem() OS {
	return mem.New()
}
