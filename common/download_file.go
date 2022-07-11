package common

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/spf13/viper"
)

type FileIdRequest struct {
	Ok     bool `json:"ok"`
	Result struct {
		FileID       string `json:"file_id"`
		FileUniqueID string `json:"file_unique_id"`
		FileSize     int    `json:"file_size"`
		FilePath     string `json:"file_path"`
	} `json:"result"`
}

// DownloadFile -> downloading file then send to telegram
func DownloadFile(filepath, url string) error {

	timeout := 30000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	res, err := client.Get(url, nil)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	var responseTeleFile FileIdRequest
	err = json.Unmarshal(body, &responseTeleFile)
	if err != nil {
		return err
	}

	// Get the file data
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", viper.GetString("teleterm.telegram_token"), responseTeleFile.Result.FilePath))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	return err
}
