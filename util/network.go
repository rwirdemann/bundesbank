package util

import (
	"os"
)

func GetHostname() string {
	if hostname, err := os.Hostname(); err == nil && hostname == "golem" {
		return "localhost"
	}

	return "94.130.79.196"
}
