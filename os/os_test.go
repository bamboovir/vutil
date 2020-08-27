package os

import (
	"testing"

	"dev.azure.com/humana/pie/lib/os/mem"
	"dev.azure.com/humana/pie/lib/os/os"
	"github.com/stretchr/testify/assert"
)

func TestNewOS(t *testing.T) {
	assert.NotNil(t, NewOS())
	assert.IsType(t, &os.OS{}, NewOS())
}

func TestNewMem(t *testing.T) {
	assert.NotNil(t, NewMem())
	assert.IsType(t, &mem.OS{}, NewMem())
}
