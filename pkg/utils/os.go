package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"tools/pkg/config"
)

func Mkdir(dirPath string) {
	os.Mkdir(dirPath, 0777)
}

func FileExist(file string) bool {
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ShowFileExist(files ...string) bool {
	for _, filePath := range files {
		if !FileExist(filePath) {
			msg := fmt.Sprintf(`[ %s ] file not exist !!!`, filePath)
			fmt.Println(msg)
			return false
		}
	}
	return true
}

func GetFileName(localPath string) string {
	return filepath.Base(localPath)
}

func GetFileBaseName(localPath string) string {
	fileName := GetFileName(localPath)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func GetFullPath(dirPath, fileName string) string {
	filepath := filepath.Join(dirPath, fileName)
	return strings.Replace(filepath, `\`, `/`, -1)
}

func CreateBlankFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	printError(err)
	defer file.Close()
}

func RefreshFileTime(filePath string) {
	now := time.Now()
	err := os.Chtimes(filePath, now, now)
	printError(err)
}

func RunCmd(cmd string) string {
	var bash, args = `bash`, `-c`
	if config.IsWin {
		bash, args = `cmd`, `/c`
	}
	output, _ := exec.Command(bash, args, cmd).CombinedOutput()
	return string(output)
}

func PrintError(err error, msgs ...string) {
	if err != nil {
		for _, msg := range msgs {
			fmt.Println(msg)
		}
		os.Exit(1)
	}
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
