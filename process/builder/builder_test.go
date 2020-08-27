package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder_Copy(t *testing.T) {
	type fields struct {
		Name    string
		Args    []string
		WorkDir string
		Env     map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		want   *Builder
	}{
		{
			name: "copy",
			fields: fields{
				Name:    "a",
				Args:    []string{"a", "b"},
				WorkDir: "/abc",
				Env: map[string]string{
					"a": "b",
					"c": "d",
				},
			},
			want: &Builder{
				Name:    "a",
				Args:    []string{"a", "b"},
				WorkDir: "/abc",
				Env: map[string]string{
					"a": "b",
					"c": "d",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Builder{
				Name:    tt.fields.Name,
				Args:    tt.fields.Args,
				WorkDir: tt.fields.WorkDir,
				Env:     tt.fields.Env,
			}

			assert.NotSame(t, tt.want, b.Copy())
			assert.Equal(t, tt.want, b.Copy())
		})
	}
}

func TestNew(t *testing.T) {
	b := New("terraform")
	assert.NotNil(t, b)
	assert.IsType(t, &Builder{}, b)
	assert.Equal(t, "terraform", b.Name)
}

func TestBuilder_WithName(t *testing.T) {
	b := &Builder{
		Name: "",
	}

	b.WithName("a")

	assert.Equal(t, "a", b.Name)
}

func TestBuilder_WithArgs(t *testing.T) {
	b := &Builder{
		Args: []string{},
	}

	b.WithArgs([]string{"a", "b"})

	assert.Equal(t, []string{"a", "b"}, b.Args)
}

func TestBuilder_AppendArgs(t *testing.T) {
	b := &Builder{
		Args: []string{},
	}

	b.AppendArgs([]string{"a", "b"})

	assert.Equal(t, []string{"a", "b"}, b.Args)

	b.AppendArgs([]string{"c", "d"})

	assert.Equal(t, []string{"a", "b", "c", "d"}, b.Args)
}

func TestBuilder_WithWorkDir(t *testing.T) {
	b := &Builder{
		WorkDir: "/",
	}

	b.WithWorkDir("/usr/local/")

	assert.Equal(t, "/usr/local/", b.WorkDir)
}

func TestBuilder_WithEnv(t *testing.T) {
	b := &Builder{
		Env: map[string]string{},
	}

	ori := map[string]string{
		"a": "b",
		"c": "d",
	}

	b.WithEnv(ori)

	assert.Equal(t, ori, b.Env)
}

func TestProcessBuilder_WithEnvAppend(t *testing.T) {
	oriA := map[string]string{
		"a": "b",
		"c": "d",
	}

	oriB := map[string]string{
		"a": "d",
		"e": "f",
	}

	oriC := map[string]string{
		"a": "d",
		"c": "d",
		"e": "f",
	}

	b := &Builder{
		Env: oriA,
	}

	b.WithEnvAppend(oriB)

	assert.Equal(t, oriC, b.Env)
}
