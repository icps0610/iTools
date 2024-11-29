package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"tools/pkg/cmdutils"
	"tools/pkg/utils"
)

func main() {
	appVersion := `1.0`
	argsCount := 2
	description := `This program removes EXIF data from image by reconstructing them.`
	useage := `Usage: rmEXIF [inputPath] [outputPath]`

	var args []string

	// 設定參數
	{
		cmdutils.Set(useage, description)
		args = cmdutils.SetArgs(argsCount, appVersion)
	}

	// 主程式
	{
		inputPath, outputPath := args[0], args[1]
		// 檢查是否存在
		if !utils.ShowFileExist(inputPath) {
			return
		}

		ext := filepath.Ext(inputPath)

		ioReader, _ := os.Open(inputPath)
		img, _, err := image.Decode(ioReader)
		printError(err)

		file, err := os.Create(outputPath)
		printError(err)
		defer file.Close()

		if ext == `.png` {
			err = png.Encode(file, img)
		} else if ext == `.jpeg` || ext == `.jpg` {
			err = jpeg.Encode(file, img, nil)
		}

		printError(err)
	}
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
