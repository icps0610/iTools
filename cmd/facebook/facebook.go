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
	description := `This program posts a message to a Facebook page.`
	useage := `Usage: facebook --token [token] (--id [id] | --postID [postID]) --msg [message]

  -t, --token              token (read from the environment variable
                           [FbToken] if not provided).
  -i, --id                 ID (optional, required if postID is not provided).
  -p, --postID             Post ID (optional, required if ID is not provided).
  -m, --msg                Message to send (required).

Example:
    facebook -t %FbToken% -i %FbFansID% -m test
    facebook -t %FbToken% -p [postID] -m test
    facebook -t %FbToken% -i %FbFansID% --list`

	var token, id, postID, message string
	var list bool
	// 設定參數
	{
		cmdutils.Set(useage, description)

		flag.StringVar(&token, "t", "", "token")
		flag.StringVar(&token, "token", "", "token")

		flag.StringVar(&id, "i", "", "id")
		flag.StringVar(&id, "id", "", "id")

		flag.StringVar(&postID, "p", "", "postID")
		flag.StringVar(&postID, "postID", "", "postID")

		flag.StringVar(&message, "m", "", "message")
		flag.StringVar(&message, "msg", "", "message")

		flag.BoolVar(&list, "list", false, "list")

		cmdutils.SetArgs(argsCount, appVersion)
	}

	// 讀取系統變數
	{
		envFbToken := os.Getenv("FbToken")
		if envFbToken != "" {
			token = envFbToken
		}
	}

	// 主程式
	{

		// 若是查詢
		if list {
			url := fmt.Sprintf(`https://graph.facebook.com/v21.0/%s/posts?fields=id&limit=10&access_token=%s`, id, token)

			respone := utils.SendGet(url)

			for _, id := range utils.Scans(respone, `(\d+_\d+)`) {
				fmt.Println(id)
			}

			os.Exit(0)
		}

		// 發表
		// required token id message
		if token == "" || message == "" {
			cmdutils.ShowHelp()
		}

		var url string
		if id != "" && postID == "" {
			url = fmt.Sprintf(`https://graph.facebook.com/%s/feed`, id)
		} else if id == "" && postID != "" {
			url = fmt.Sprintf(`https://graph.facebook.com/v21.0/%s/comments`, postID)
		}

		formData := map[string]string{
			"message":      message,
			"access_token": token,
		}

		respone := utils.SendPost(url, "multipart/form-data", formData)
		fmt.Println(respone)
	}
}
