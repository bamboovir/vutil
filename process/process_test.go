package process

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"dev.azure.com/humana/pie/lib/process/mem"
	"dev.azure.com/humana/pie/lib/process/os"
)

func TestOS(t *testing.T) {
	pc := OS()
	assert.NotNil(t, pc)
	assert.IsType(t, &os.Processor{}, pc)
}

func TestMem(t *testing.T) {
	pc := Mem()
	assert.NotNil(t, pc)
	assert.IsType(t, &mem.Processor{}, pc)
}
