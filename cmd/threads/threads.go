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
	description := `This program posts messages to Threads.`
	useage := `Usage: threads --token [token] --id [id] --msg [message]

  -t, --token              token (required).
  -i, --id                 ID  (required)
  -m, --msg                Message to send (required).`

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
		envThreadsToken := os.Getenv("ThreadsToken")
		if envThreadsToken != "" {
			token = envThreadsToken
		}
		envThreadsUserID := os.Getenv("ThreadsUserID")
		if envThreadsUserID != "" {
			id = envThreadsUserID
		}
	}

	// 主程式
	{

		// required token id message
		if token == "" || id == "" || message == "" {
			cmdutils.ShowHelp()
		}

		Threads_URL := fmt.Sprintf(`https://graph.threads.net/v1.0/%s/`, id)

		var containerID string
		{
			url := Threads_URL + `threads`

			jsonData := map[string]interface{}{
				"access_token": token,
				"media_type":   "TEXT",
				"text":         message,
			}

			respone := utils.SendPost(url, "application/json", jsonData)

			containerID = utils.Scan(respone, `{"id":"(\d+)"}`, 1)
		}

		var postID string
		{
			url := Threads_URL + `threads_publish`

			jsonData := map[string]interface{}{
				"access_token": token,
				"creation_id":  containerID,
			}

			respone := utils.SendPost(url, "application/json", jsonData)

			postID = utils.Scan(respone, `{"id":"(\d+)"}`, 1)
		}

		fmt.Println(containerID, postID)

	}
}
