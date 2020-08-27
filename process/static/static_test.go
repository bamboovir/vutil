package static

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyStringMap(t *testing.T) {
	ori := map[string]string{
		"a": "b",
		"c": "d",
	}

	rst := CopyStringMap(ori)

	assert.NotSame(t, ori, rst)
	assert.Equal(t, ori, rst)
}

func TestMergeEnvMap(t *testing.T) {
	rst := MergeEnvMap(
		map[string]string{"a": "b"},
		map[string]string{"a": "c"},
		map[string]string{"c": "d"},
		map[string]string{"e": "f"},
	)

	assert.Equal(t, map[string]string{
		"a": "c",
		"c": "d",
		"e": "f",
	}, rst)
}

func TestEnvMapToStringSlice(t *testing.T) {
	rst := EnvMapToStringSlice(map[string]string{
		"a": "b",
		"c": "d",
	})

	sort.Strings(rst)

	assert.Equal(t, []string{"a=b", "c=d"}, rst)
}
