package static

import (
	"context"
	"io"
)

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

// EnvMapToStringSlice could convert a environment map into a string
func EnvMapToStringSlice(env map[string]string) []string {
	envList := make([]string, 0, 0)

	for k, v := range env {
		envList = append(envList, k+"="+v)
	}
	return envList
}

// CancellableCopy define
func CancellableCopy(ctx context.Context, dst io.Writer, src io.Reader) (written int64, err error) {
	size := 32 * 1024
	buf := make([]byte, size)

	for {
		select {
		case <-ctx.Done():
			return written, nil
		default:
			nr, er := src.Read(buf)
			if nr > 0 {
				nw, ew := dst.Write(buf[0:nr])
				if nw > 0 {
					written += int64(nw)
				}
				if ew != nil {
					err = ew
					return written, err
				}
				if nr != nw {
					err = io.ErrShortWrite
					return written, err
				}
			}
			if er != nil {
				if er != io.EOF {
					err = er
				}
				return written, err
			}
		}
	}

}
