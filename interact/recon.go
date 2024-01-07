package interact

import (
	"runtime"
)

// Recon functions

func GetOsName() string {
	return runtime.GOOS
}
