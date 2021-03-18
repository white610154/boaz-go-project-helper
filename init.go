package main

import (
	"os"
	"strings"
)

func get_gopath() string {
	for _, l := range os.Environ() {
		pair := strings.SplitN(l, "=", 2)
		if pair[0] == "GOPATH" {
			return pair[1]
		}
	}
	return ""
}
