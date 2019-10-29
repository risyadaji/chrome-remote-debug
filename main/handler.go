package main

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
)

// get post from crawler
// decode
// define bank
// get statement from selected bank
// upload to crawler uploader
// upload to upload endpoint (set env for avoid rebuild if there something changes)

func getStatemenHandler(w http.ResponseWriter, r *http.Request) {
	var p Payload

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := UploadFile(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Success"))
	return
}

// UploadFile ...
func UploadFile(p Payload) error {
	path := fmt.Sprintf("%s", p.FileName)
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

	writer.WriteField("accountName", p.AccountName)

	if err := writer.Close(); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", p.UploadLink, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-Api-Key", p.APIKey)
	req.Header.Set("Cache-Control", "no-cache")

	client := &http.Client{}
	if _, err := client.Do(req); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
