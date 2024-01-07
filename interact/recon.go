package go4it

import (
	"runtime"
)

// Recon functions

func GetOsName() string {
	return runtime.GOOS
}
