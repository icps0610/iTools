package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	RootDirPath = getRootDirPath()
	TempDirPath = getTempDirPath()
	IsWin       = runtime.GOOS == "windows"
)

func getRootDirPath() string {
	path, err := os.Executable()
	printError(err)
	if strings.HasPrefix(strings.ToLower(filepath.Dir(path)), strings.ToLower(os.TempDir())) {
		path, err = os.Getwd()
		printError(err)
		return path + `\`
	}
	path = filepath.Dir(path)
	if IsWin {
		return path + `\`
	}
	return path + `/`
}

func getTempDirPath() string {
	if IsWin {
		return `z:\matsu`
	}
	return `/tmp/matsu`
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
