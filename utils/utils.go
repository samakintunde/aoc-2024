package utils

import (
	"os"
	"strings"
)

func ReadInput(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	contents := strings.TrimSpace(string(bytes))
	return contents
}
