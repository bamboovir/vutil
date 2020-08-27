package os

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, New())
	assert.IsType(t, &Env{}, New())
}

func SetupOSEnvCase() {
	os.Clearenv()
	os.Setenv("a", "b")
}

func TestEnviron(t *testing.T) {
	e := New()
	SetupOSEnvCase()

	rst := e.Environ()
	assert.Equal(t, []string{"a=b"}, rst)
}

func TestGetenv(t *testing.T) {
	e := New()
	SetupOSEnvCase()

	rst := e.Getenv("a")
	assert.Equal(t, "b", rst)

	rst = e.Getenv("c")
	assert.Equal(t, "", rst)
}

func TestLookupEnv(t *testing.T) {
	e := New()
	SetupOSEnvCase()
	rst, exist := e.LookupEnv("a")
	assert.True(t, exist)
	assert.Equal(t, "b", rst)

	rst, exist = e.LookupEnv("c")
	assert.False(t, exist)
	assert.Equal(t, "", rst)
}

func TestClearenv(t *testing.T) {
	e := New()
	SetupOSEnvCase()
	rst := e.Environ()
	assert.Equal(t, []string{"a=b"}, rst)

	e.Clearenv()
	assert.Equal(t, []string{}, e.Environ())
}

func TestSetenv(t *testing.T) {
	e := New()

	e.Clearenv()
	assert.Equal(t, []string{}, e.Environ())

	e.Setenv("a", "b")
	assert.Equal(t, []string{"a=b"}, e.Environ())
}
