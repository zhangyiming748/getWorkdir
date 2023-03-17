package getWorkdir

import (
	"path"
	"runtime"
)

func GetRoot() string {
	_, filename, _, _ := runtime.Caller(0)

	return path.Dir(filename)
}
