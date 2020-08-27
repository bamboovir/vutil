package environment

import (
	"testing"

	"dev.azure.com/humana/pie/lib/environment/mem"
	"dev.azure.com/humana/pie/lib/environment/os"
	"github.com/stretchr/testify/assert"
)

func TestOS(t *testing.T) {
	assert.NotNil(t, OS())
	assert.IsType(t, &os.Env{}, OS())
}

func TestMem(t *testing.T) {
	assert.NotNil(t, Mem([]string{}))
	assert.IsType(t, &mem.Env{}, Mem([]string{}))
}

func TestMemFromMap(t *testing.T) {
	assert.NotNil(t, MemFromMap(map[string]string{}))
	assert.IsType(t, &mem.Env{}, MemFromMap(map[string]string{}))
}
