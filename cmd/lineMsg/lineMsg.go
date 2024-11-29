package main

import (
	"flag"
	"fmt"
	"os"

	"tools/pkg/cmdutils"
	"tools/pkg/utils"
)

func main() {
	appVersion := `1.0`
	argsCount := 0
	description := `This program use LINE Messaging API to send message`
	useage := `Usage: lineMsg --token [token] --id [id] --msg [message]

  -t, --token              token (read from the environment variable
                           [LineMsgToken] if not provided).
  -i, --id                 ID  (required)
  -m, --msg                Message to send (required).
  
Example: 
    lineMsg -m test`

	var token, id, message string

	// 設定參數
	{
		cmdutils.Set(useage, description)

		flag.StringVar(&token, "t", "", "token")
		flag.StringVar(&token, "token", "", "token")

		flag.StringVar(&id, "i", "", "id")
		flag.StringVar(&id, "id", "", "id")

		flag.StringVar(&message, "m", "", "message")
		flag.StringVar(&message, "msg", "", "message")

		cmdutils.SetArgs(argsCount, appVersion)
	}

	// 讀取系統變數
	{
		envLineMsgToken := os.Getenv("LineMsgToken")
		if envLineMsgToken != "" {
			token = envLineMsgToken
		}
	}

	// 主程式
	{
		// required token id message
		if token == "" || id == "" || message == "" {
			cmdutils.ShowHelp()
		}

		url := "https://api.line.me/v2/bot/message/push"

		jsonData := map[string]interface{}{
			"to": id,
			"messages": []map[string]string{
				{
					"type": "text",
					"text": message,
				},
			},
		}

		response := utils.SendPost(url, "application/json", jsonData, token)

		fmt.Println(response)
	}
}
