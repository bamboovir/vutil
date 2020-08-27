package mem

// import (
// 	"bytes"
// 	"testing"

// 	"dev.azure.com/humana/pie/lib/process/builder"
// 	"github.com/stretchr/testify/assert"
// )

// func TestWithStdout(t *testing.T) {
// 	b := builder.New("terraform")
// 	p := New(*b)

// 	w := bytes.NewBuffer(make([]byte, 0))
// 	pA := p.WithStdout(w)

// 	assert.Same(t, pA, p)
// 	assert.Same(t, w, p.stdout)

// }

// func TestWithStderr(t *testing.T) {
// 	b := builder.New("terraform")
// 	p := New(*b)

// 	w := bytes.NewBuffer(make([]byte, 0))
// 	pA := p.WithStderr(w)

// 	assert.Same(t, pA, p)
// 	assert.Same(t, w, p.stderr)
// }

// func TestWithStdin(t *testing.T) {
// 	b := builder.New("terraform")
// 	p := New(*b)

// 	r := bytes.NewBuffer(make([]byte, 0))
// 	pA := p.WithStdin(r)

// 	assert.Same(t, pA, p)
// 	assert.Same(t, r, p.stdin)
// }

// func TestName(t *testing.T) {
// 	b := builder.New("terraform")
// 	p := New(*b)

// 	assert.Equal(t, "terraform", p.Name())
// }

// func TestTerminate(t *testing.T) {

// }

// func TestExec(t *testing.T) {

// }
