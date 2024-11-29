package main

import (
	"flag"
	"fmt"
	"math/rand"

	"tools/pkg/cmdutils"
)

var defaultChar = `23456789ABCDEFGHJKMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz!@#$%^&*`

func main() {
	appVersion := `1.0`
	argsCount := 0
	description := fmt.Sprintf(`This program generates random password. default characters:
%s`, defaultChar)
	useage := `Usage: passwd --char [words] --len [number]

  -c, --char               Customize the character set used for generating passwords.
  -l, --len                Specify the length of the generated password.`

	var char string
	var length int

	// 設定參數
	{
		cmdutils.Set(useage, description)

		flag.StringVar(&char, "c", defaultChar, "character")
		flag.StringVar(&char, "char", defaultChar, "character")

		flag.IntVar(&length, "l", 10, "length")
		flag.IntVar(&length, "len", 10, "length")

		cmdutils.SetArgs(argsCount, appVersion)
	}

	// 主程式
	{
		count := len(char)

		var result string
		for len(result) < length {
			result += string(char[rand.Intn(count)])
		}
		fmt.Println(result)
	}
}
