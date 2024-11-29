package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"time"
)

func SendGet(url string) string {
	res, err := http.Get(url)
	printError(err)
	if res == nil {
		fmt.Println(url, "no response")
		return ""
	} else {
		defer res.Body.Close()
	}

	res.Header.Set(`User-Agent`, `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36`)
	body, err := io.ReadAll(res.Body)
	printError(err)

	// WriteFile(string(body), `z:\log.html`)

	return string(body)
}

func SendPost(url, contentType string, data interface{}, apiKeys ...string) string {
	var req *http.Request
	var err error

	if contentType == "application/json" {
		// 將資料序列化為 JSON 格式
		jsonData, jsonErr := json.Marshal(data)
		printError(jsonErr)

		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

	} else if contentType == "multipart/form-data" {
		// 處理 multipart/form-data
		form := new(bytes.Buffer)
		writer := multipart.NewWriter(form)

		// 檢查資料是否包含文字和檔案字段
		if fields, ok := data.(map[string]interface{}); ok {
			for key, value := range fields {
				switch v := value.(type) {
				case string:
					// 處理文字字段
					formField, formErr := writer.CreateFormField(key)
					printError(formErr)
					_, writeErr := formField.Write([]byte(v))
					printError(writeErr)

				case *os.File:
					// 處理檔案字段，字段名稱需為 "photo"
					if key != "photo" {
						fmt.Println("Unsupported file field key:", key)
						continue
					}
					part, formErr := writer.CreateFormFile(key, v.Name())
					printError(formErr)
					_, copyErr := io.Copy(part, v)
					printError(copyErr)
					v.Close() // 確保檔案處理結束後關閉
				default:
					fmt.Printf("Unsupported field type for key '%s': %T\n", key, v)
				}
			}
		}

		writer.Close()

		req, err = http.NewRequest("POST", url, form)
		req.Header.Set("Content-Type", writer.FormDataContentType())
	} else {
		// 若 contentType 無法識別，回傳錯誤
		printError(fmt.Errorf("unsupported content type: %s", contentType))
		return ""
	}

	if len(apiKeys) > 0 {
		req.Header.Set("Authorization", "Bearer "+apiKeys[0])
	}

	printError(err)

	client := &http.Client{}
	res, resErr := client.Do(req)
	printError(resErr)
	if res == nil {
		fmt.Println(url, "no response")
		return ""
	} else {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	printError(readErr)

	return string(body)
}

func CheckPort(ip, port string, wait int) error {
	address := fmt.Sprintf("%s:%s", ip, port)
	_, err := net.DialTimeout("tcp", address, time.Duration(wait)*time.Millisecond)
	return err
}
