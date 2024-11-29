package utils

import (
	"os"
	"strings"
)

func WriteFile(content, path string) {
	err := os.WriteFile(path, []byte(content), 0777)
	printError(err)
}

func ReadLines(path string) []string {
	bytes, err := os.ReadFile(path)
	printError(err)
	var contents []string
	for _, line := range strings.Split(string(bytes), "\n") {
		contents = append(contents, strings.TrimSpace(line))
	}
	return contents
}
