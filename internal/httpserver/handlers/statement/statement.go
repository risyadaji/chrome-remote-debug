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

// Payload represents params for upload
type Payload struct {
	AccountName string `json:"accountName"`
	PostDate    string `json:"postDate"`
	FileName    string `json:"fileName"`
	UploadLink  string `json:"uploadLink"`
	APIKey      string `json:"apiKey"`
}

// UploadResponse represent responses data after upload
type UploadResponse struct {
	TotalStatement int `json:"totalStatement"`
}

// GetHandler ..
func GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := Payload{}

		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			return
		}

		res, err := uploadFile(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response.JSON(w, http.StatusOK, res)
		return
	}
}

// UploadFile ...
func uploadFile(p Payload) (*UploadResponse, error) {
	path := config.ChromeDownloadPath()
	path += p.FileName
	log.Println(path)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("data", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, err
	}

	writer.WriteField("accountName", p.AccountName)
	writer.WriteField("postDate", p.PostDate)
	if err := writer.Close(); err != nil {
		return nil, err
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
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf(res.Status)
	}

	var response UploadResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	log.Println(response)

	// Delete file
	if err := os.Remove(path); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &response, nil
}
