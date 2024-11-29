package main

import (
	"flag"
	"fmt"

	"tools/pkg/cmdutils"
	"tools/pkg/utils"
)

func main() {
	appVersion := `1.0`
	argsCount := 0
	description := `This program displays the current time or date. If no options are specified, the default output 
is the current month, day, hour, and minute in the format MMDDHHMM (e.g., 06100030)`
	useage := `Usage: now [options]

  -f, --full         Display the full timestamp (e.g., 2024-01-01-00:00:00). Optional.
  -d, --day          Display the current date (e.g., 20240101). Optional.

  -h, --help         Show this help message and exit.
  -v, --version      Display version information and exit.`

	var full, day bool

	// 設定參數
	{
		cmdutils.Set(useage, description)

		flag.BoolVar(&full, "f", false, "Display the full timestamp")
		flag.BoolVar(&full, "full", false, "Display the full timestamp")

		flag.BoolVar(&day, "d", false, "Display the current date")
		flag.BoolVar(&day, "day", false, "Display the current date")

		cmdutils.SetArgs(argsCount, appVersion)
	}

	// 主程式
	{
		now := utils.TimeNow()

		var msg string
		if full {
			msg = fmt.Sprintf(`%v-%02v-%02v-%02v:%02v:%02v`, now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
		} else if day {
			msg = fmt.Sprintf(`%v%02v%02v`, now.Year(), int(now.Month()), now.Day())
		} else {
			msg = fmt.Sprintf(`%02v%02v%02v%02v`, int(now.Month()), now.Day(), now.Hour(), now.Minute())
		}

		fmt.Println(msg)
	}

}
