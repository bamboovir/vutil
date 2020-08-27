package os

import "os"

// Env define
type Env struct{}

// New define
func New() *Env {
	return &Env{}
}

// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
// To distinguish between an empty value and an unset value, use LookupEnv.
func (e *Env) Getenv(key string) string {
	return os.Getenv(key)
}

// LookupEnv retrieves the value of the environment variable named by the key.
// If the variable is present in the environment the value (which may be empty)
// is returned and the boolean is true.
// Otherwise the returned value will be empty and the boolean will be false.
func (e *Env) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

// Setenv sets the value of the environment variable named by the key.
// It never returns an error, if any.
func (e *Env) Setenv(key, value string) error {
	return os.Setenv(key, value)
}

// Environ returns a copy of strings representing the environment,
// in the form "key=value".
func (e *Env) Environ() []string {
	return os.Environ()
}

// Clearenv deletes all environment variables.
func (e *Env) Clearenv() {
	os.Clearenv()
}
