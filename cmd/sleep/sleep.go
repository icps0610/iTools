package main

import (
	"fmt"

	"tools/pkg/cmdutils"
	"tools/pkg/utils"
)

func main() {
	appVersion := `1.0`
	argsCount := 1
	description := `This program pauses execution for a specified number of seconds.`
	useage := `Usage: sleep [number]`

	var args []string

	// 設定參數
	{
		cmdutils.Set(useage, description)
		args = cmdutils.SetArgs(argsCount, appVersion)
	}
	// 主程式
	{
		seconds := utils.To_i(args[0]) * 1000
		fmt.Println(seconds, `sec`)

		utils.Sleep(seconds)
	}

}
