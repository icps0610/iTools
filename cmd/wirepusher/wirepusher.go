package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"tools/pkg/cmdutils"
	"tools/pkg/utils"
)

func main() {
	appVersion := `1.0`
	argsCount := 0
	description := `This program use WirePusher to send message`
	useage := `Usage: wirepusher --token [token] --type [type] --title [title] --msg [message]

  -t, --token              Wirepusher token (read from the environment variable
                           [WirePusherToken] if not provided).
      --type               type (required).
      --title              title (required).
  -m, --msg                Message to send (required).
  -a, --action             URL (optional).`

	var token, group, title, msg, action string

	// 設定參數
	{
		cmdutils.Set(useage, description)

		flag.StringVar(&token, "t", "", "title")
		flag.StringVar(&token, "token", "", "token")

		flag.StringVar(&group, "type", "", "type")

		flag.StringVar(&title, "title", "", "title")

		flag.StringVar(&msg, "m", "", "msg")
		flag.StringVar(&msg, "msg", "", "msg")

		flag.StringVar(&action, "a", "", "action")
		flag.StringVar(&action, "action", "", "action")

		cmdutils.SetArgs(argsCount, appVersion)
	}

	// 讀取系統變數
	{
		envWirePusherToken := os.Getenv("WirePusherToken")
		if envWirePusherToken != "" {
			token = envWirePusherToken
		}
	}

	// 主程式
	{
		// required token msg
		if token == "" || group == "" || title == "" || msg == "" {
			cmdutils.ShowHelp()
		}

		msg = url.QueryEscape(msg)

		if action != "" {
			msg += fmt.Sprintf(`&action=%s`, action)
		}

		url := fmt.Sprintf(`https://wirepusher.com/send?id=%s&type=%s&title=%s&message=%s`, token, group, title, msg)

		fmt.Println(url)
		utils.SendGet(url)
	}

}
