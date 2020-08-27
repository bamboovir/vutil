package mem

import (
	"dev.azure.com/humana/pie/lib/environment/static"
)

// Env define
type Env struct {
	m map[string]string
}

// New define
func New(envs []string) *Env {
	set := &Env{}
	set.m = static.EnvStringSliceToMap(envs)
	return set
}

// Environ returns a copy of strings representing the environment,
// in the form "key=value".
func (e *Env) Environ() []string {
	return static.EnvMapToStringSlice(e.m)
}

// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
// To distinguish between an empty value and an unset value, use LookupEnv.
func (e *Env) Getenv(key string) string {
	val, ok := e.m[key]
	if !ok {
		return ""
	}
	return val
}

// LookupEnv retrieves the value of the environment variable named by the key.
// If the variable is present in the environment the value (which may be empty)
// is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will be false.
func (e *Env) LookupEnv(key string) (string, bool) {
	val, ok := e.m[key]
	return val, ok
}

// Clearenv deletes all environment variables.
func (e *Env) Clearenv() {
	e.m = make(map[string]string)
}

// Setenv sets the value of the environment variable named by the key.
// It never returns an error, if any.
func (e *Env) Setenv(key string, value string) error {
	e.m[key] = value
	return nil
}
