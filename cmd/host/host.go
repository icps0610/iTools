package main

import (
	"fmt"
	"net"

	"tools/pkg/cmdutils"
	"tools/pkg/utils"
)

func main() {
	appVersion := `1.0`
	argsCount := 1
	description := `This program shows the target hostname.`
	useage := `Usage: host [ip]`

	var args []string

	// 設定參數
	{
		cmdutils.Set(useage, description)
		args = cmdutils.SetArgs(argsCount, appVersion)
	}

	// 主程式
	{
		ip := args[0]
		names, err := net.LookupAddr(ip)
		utils.PrintError(err)

		for _, name := range names {
			fmt.Println(name)
		}
	}
}
