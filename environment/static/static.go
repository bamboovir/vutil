package static

import (
	"strings"
)

// EnvMapToStringSlice could convert a environment map into a string
func EnvMapToStringSlice(env map[string]string) []string {
	envList := make([]string, 0, 0)

	for k, v := range env {
		envList = append(envList, k+"="+v)
	}
	return envList
}

// EnvStringSliceToMap could convert a environment map into a map
func EnvStringSliceToMap(envs []string) map[string]string {
	envList := make(map[string]string, len(envs))

	for _, s := range envs {
		kv := strings.SplitN(s, "=", 2)
		envList[kv[0]] = kv[1]
	}

	return envList
}

// CopyStringMap define
func CopyStringMap(original map[string]string) map[string]string {
	copyMap := make(map[string]string, len(original))
	for key, value := range original {
		copyMap[key] = value
	}
	return copyMap
}

// MergeEnvMap could merge envs into once
// It will overwrite first appeared Env pair
func MergeEnvMap(envs ...map[string]string) map[string]string {
	mergedEnvMap := map[string]string{}

	for _, env := range envs {
		for k, v := range env {
			mergedEnvMap[k] = v
		}
	}

	return mergedEnvMap
}
