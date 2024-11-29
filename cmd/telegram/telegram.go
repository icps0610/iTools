package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"tools/pkg/cmdutils"
	"tools/pkg/utils"
)

func main() {
	appVersion := `1.0`
	argsCount := 0
	description := `This program uses the Telegram API to send messages.`
	useage := `Usage: telegram --token [token] --id [id] --msg [message]

  -t, --token              Bot token (reads from the environment variable
                           [TelegramToken] if not provided).
  -i, --id                 Chat ID (reads from the environment variable
                           [TelegramChatID] if not provided).
      --title              Title for the message (optional).
  -m, --msg                Text message to send (required).

      --markdown           Format the message using Markdown (optional).
      --preview            Disable web page preview in the message (optional).
  -p, --photoPath          Path to a photo to send (optional).
  
 Example: 
    telegram -m test`

	var token, id, title, message, photoPath string

	var markdwon, preview bool
	var markdwonStr, previewStr = "", "false"
	// 設定參數
	{
		cmdutils.Set(useage, description)

		flag.StringVar(&token, "t", "", "token")
		flag.StringVar(&token, "token", "", "token")

		flag.StringVar(&id, "i", "", "id")
		flag.StringVar(&id, "id", "", "id")

		flag.StringVar(&title, "title", "", "title")

		flag.StringVar(&message, "m", "", "message")
		flag.StringVar(&message, "msg", "", "message")

		flag.StringVar(&photoPath, "p", "", "photoPath")
		flag.StringVar(&photoPath, "photoPath", "", "photoPath")

		flag.BoolVar(&markdwon, "markdown", false, "markdwon")
		flag.BoolVar(&preview, "preview", false, "disable_web_page_preview")

		cmdutils.SetArgs(argsCount, appVersion)
	}

	// 讀取系統變數
	{
		envTelegramToken := os.Getenv("TelegramToken")
		if envTelegramToken != "" {
			token = envTelegramToken
		}

		envTelegramChatID := os.Getenv("TelegramChatID")
		if envTelegramChatID != "" {
			id = envTelegramChatID
		}
	}

	// 主程式
	{
		// required token id message
		if token == "" || id == "" {
			cmdutils.ShowHelp()
		}

		var apiURL = fmt.Sprintf(`https://api.telegram.org/bot%s/`, token)

		if photoPath != "" {
			url := apiURL + "sendPhoto"

			// 準備檔案
			file, err := os.Open(photoPath)
			printError(err)
			defer file.Close()

			// 準備表單資料
			data := map[string]interface{}{
				"chat_id": id,
				"photo":   file,
			}

			// 發送 POST 請求
			response := utils.SendPost(url, "multipart/form-data", data)
			fmt.Println("Response:", response)

			os.Exit(0)

		} else if message != "" {
			u := apiURL + "sendMessage"

			if markdwon {
				markdwonStr = "MarkdownV2"

				title = reservedWords(title)
				message = reservedWords(message)

			}

			if title != "" {
				message = fmt.Sprintf("%s\n%s", title, message)
			}

			if preview == true {
				previewStr = "true"
			}

			jsonData := map[string]interface{}{
				"chat_id":    id,
				"text":       message,
				"parse_mode": markdwonStr,
				// 取消網頁預覽
				"disable_web_page_preview": previewStr,
			}

			response := utils.SendPost(u, "application/json", jsonData, token)
			fmt.Println(response)

			os.Exit(0)
		}

		cmdutils.ShowHelp()
	}
}

func reservedWords(str string) string {
	str = strings.ReplaceAll(str, `.`, `\.`)
	str = strings.ReplaceAll(str, `=`, `\=`)
	str = strings.ReplaceAll(str, `-`, `\-`)
	return str
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
