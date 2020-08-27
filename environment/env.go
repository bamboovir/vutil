package environment

import (
	"dev.azure.com/humana/pie/lib/environment/mem"
	"dev.azure.com/humana/pie/lib/environment/os"
	"dev.azure.com/humana/pie/lib/environment/static"
)

// EnvSet define
type EnvSet interface {
	Environ() []string
	Getenv(key string) string
	LookupEnv(key string) (string, bool)
	Setenv(key, value string) error
	Clearenv()
}

// OS define
func OS() EnvSet {
	return os.New()
}

// Mem define
func Mem(envs []string) EnvSet {
	return mem.New(envs)
}

// MemFromMap define
func MemFromMap(envs map[string]string) EnvSet {
	return mem.New(static.EnvMapToStringSlice(envs))
}
