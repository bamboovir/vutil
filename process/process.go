package process

import (
	"io"

	"dev.azure.com/humana/pie/lib/process/builder"
	"dev.azure.com/humana/pie/lib/process/mem"
	"dev.azure.com/humana/pie/lib/process/os"
	"dev.azure.com/humana/pie/lib/process/types"
)

// Processor defination
type Processor interface {
	P(builder.Builder) types.Process
}

// OS Process executor define
func OS() Processor {
	return &os.Processor{}
}

// Mem Process executor define
func Mem() Processor {
	return &mem.Processor{}
}

// CombinedOutput define
// The io.Writer instance used for writing must be concurrently safe
func CombinedOutput(p types.Process, combinedOutput io.Writer) types.Process {
	return p.WithStdout(combinedOutput).WithStderr(combinedOutput)
}
