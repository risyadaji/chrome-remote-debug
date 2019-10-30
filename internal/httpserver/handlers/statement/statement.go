package statement

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/payfazz/chrome-remote-debug/internal/httpserver/response"

	"github.com/payfazz/chrome-remote-debug/config"
)

// Payload ...
type Payload struct {
	BankName    string `json:"bankName"`
	AccountName string `json:"accountName"`
	FileName    string `json:"fileName"`
	UploadLink  string `json:"uploadLink"`
	APIKey      string `json:"apiKey"`
}

// GetHandler ..
func GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := Payload{}

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			return
		}

		if err := uploadFile(p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, http.StatusOK)
		return
	}
}

// UploadFile ...
func uploadFile(p Payload) error {
	var path string
	if config.ChromeDownloadPath() != "" {
		path = config.ChromeDownloadPath()
	}
	path += p.FileName
	log.Println(path)

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("data", filepath.Base(path))
	if err != nil {
		return err
	}
	if _, err := io.Copy(part, file); err != nil {
		return err
	}
	if err := writer.WriteField("accountName", p.AccountName); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", p.UploadLink, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Api-Key", p.APIKey)
	req.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.StatusCode, res.Status)
	// Delete file
	if err := os.Remove(path); err != nil {
		log.Fatal(err)
		return err
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf(res.Status)
	}

	return nil
}
