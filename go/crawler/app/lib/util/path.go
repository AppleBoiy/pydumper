package util

import (
	"os"
	"path"
)

func GetPath(v ...string) string {
	// Get Project Dir Path
	p, err := os.Executable()
	if err != nil {
		ErrLog("Get Executable Path Error :", err)
	}

	return path.Join(append([]string{path.Dir(p)}, v...)...)
}
