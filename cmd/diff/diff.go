package main

import (
	"flag"
	"fmt"

	"tools/internal/diff"
	"tools/pkg/cmdutils"
	"tools/pkg/utils"
)

func main() {
	appVersion := `1.0`
	argsCount := 2
	deutilsion := `This program compares two files for differences.You can specify the comparison 
mode using the --mode parameter. If you do not specify the --mode parameter, the system will
automatically use the default value diff. Available mode options include:`
	useage := `Usage: diff --mode [*diff/same/all] --output [outputFullPath] [filePath1] [filePath2]

  diff:   Shows the differences between the two files (default mode).
  same:   Returns results only when the files are identical.
  all:    Displays all content, including differences and similarities

  -m, --mode               diff/same/all default: diff
  -o, --output             Generate an HTML report of the comparison results.

Example: 
    diff -o z:\test.html z:\go.mod z:\go1.mod`

	var mode, outputPath string

	var args []string

	// 設定參數
	{
		cmdutils.Set(useage, deutilsion)

		flag.StringVar(&mode, "m", "diff", "mode")
		flag.StringVar(&mode, "mode", "diff", "mode")

		flag.StringVar(&outputPath, "o", "", "Generate an HTML report of the comparison results.")
		flag.StringVar(&outputPath, "output", "", "Generate an HTML report of the comparison results.")

		args = cmdutils.SetArgs(argsCount, appVersion)

	}

	fmt.Println(outputPath)
	// 主程式
	{
		filePath1, filePath2 := args[0], args[1]
		// 檢查是否存在
		if !utils.ShowFileExist(filePath1, filePath2) {
			return
		}

		content1, content2 := utils.ReadLines(filePath1), utils.ReadLines(filePath2)
		difference, same, all := diff.Run(content1, content2)

		fmt.Printf(`Mode: `)

		switch mode {
		case "same":
			fmt.Println(mode)
			diff.Print(same)
		case "all":
			fmt.Println(mode)
			diff.Print(all)
		default:
			fmt.Println(`diff`)
			diff.Print(difference)
		}

		if outputPath != "" {
			fmt.Println()
			fmt.Println("Generate an HTML report", outputPath)

			fmt.Println(outputPath)
			content := diff.OutputHtml(filePath1, filePath2, difference, same, all)
			utils.WriteFile(content, outputPath)
		}
	}
}
