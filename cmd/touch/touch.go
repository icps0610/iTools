package main

import (
	"tools/pkg/cmdutils"
	"tools/pkg/utils"
)

func main() {
	appVersion := `1.0`
	argsCount := 1
	description := `This program updates the timestamp of a specified file 
or creates an empty file if it does not exist.`
	useage := `Usage: touch [filePath]`

	var args []string

	// 設定參數
	{
		cmdutils.Set(useage, description)
		args = cmdutils.SetArgs(argsCount, appVersion)
	}

	// 主程式
	{
		filePath := args[0]
		if utils.FileExist(filePath) {
			utils.RefreshFileTime(filePath)
		} else {
			utils.CreateBlankFile(filePath)
		}
	}
}
