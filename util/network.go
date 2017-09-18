package util

import (
	"runtime"
)

func GetHostname() string {
	if runtime.GOOS == "windows" {
		return "localhost"
	}

	if runtime.GOOS == "linux" {
		return "94.130.79.196"
	}

	return ""
}
