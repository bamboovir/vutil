package os

import "os"

// OS defination
type OS struct{}

// Executable defination
func (o *OS) Executable() (string, error) {
	return os.Executable()
}

// UserHomeDir defination
func (o *OS) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

// UserConfigDir defination
func (o *OS) UserConfigDir() (string, error) {
	return os.UserConfigDir()
}

// UserCacheDir defination
func (o *OS) UserCacheDir() (string, error) {
	return os.UserCacheDir()
}

// Getegid defination
func (o *OS) Getegid() int {
	return os.Getegid()
}

// Geteuid defination
func (o *OS) Geteuid() int {
	return os.Geteuid()
}

// Getgid defination
func (o *OS) Getgid() int {
	return os.Getgid()
}

// Getgroups defination
func (o *OS) Getgroups() ([]int, error) {
	return os.Getgroups()
}

// Getpagesize defination
func (o *OS) Getpagesize() int {
	return os.Getpagesize()
}

// Getpid defination
func (o *OS) Getpid() int {
	return os.Getpid()
}

// Getppid defination
func (o *OS) Getppid() int {
	return os.Getppid()
}

// Getuid defination
func (o *OS) Getuid() int {
	return os.Getuid()
}

// Getwd defination
func (o *OS) Getwd() (dir string, err error) {
	return os.Getwd()
}
