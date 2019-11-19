package file

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/payfazz/chrome-remote-debug/pkg/errors"

	"github.com/payfazz/chrome-remote-debug/internal/httpserver/response"
)

var (
	// ErrDirectoryFieldEmpty error when directory field is empty
	ErrDirectoryFieldEmpty = errors.NewValidationError("directory must be filled")
	// ErrFilenameFieldEmpty error when filename field is empty
	ErrFilenameFieldEmpty = errors.NewValidationError("filename must be filled")

	// ErrFileNotFound error when looking file but not found
	ErrFileNotFound = errors.NewBaseError("ErrFileNotFound", "file not found")
	// ErrDirectoryNotExist error when looking for directory but not exist
	ErrDirectoryNotExist = errors.NewBaseError("ErrDirectoryNotFound", "directory not exist")
)

// Response ..
type Response struct {
	Filename string `json:"filename"`
}

// GetFileHandler ..
func GetFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		dirname := q.Get("directory")
		if dirname == "" {
			response.WithError(w, nil, ErrDirectoryFieldEmpty)
			return
		}
		filename := q.Get("filename")
		if filename == "" {
			response.WithError(w, nil, ErrFilenameFieldEmpty)
			return
		}

		root, err := os.Getwd()
		if err != nil {
			response.WithError(w, nil, err)
			return
		}

		folderpath := path.Join(root, dirname)
		if _, err := os.Stat(folderpath); os.IsNotExist(err) {
			response.WithError(w, nil, ErrDirectoryNotExist)
			return
		}

		if _, err := os.Stat(path.Join(folderpath, filename)); os.IsNotExist(err) {
			response.WithError(w, nil, ErrFileNotFound)
			return
		}

		response.JSON(w, http.StatusOK, Response{
			Filename: filename,
		})
	}
}

// GetFilesHandler ...
func GetFilesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		dirname := q.Get("directory")
		if dirname == "" {
			response.WithError(w, nil, ErrDirectoryFieldEmpty)
			return
		}

		root, err := os.Getwd()
		if err != nil {
			response.WithError(w, nil, err)
			return
		}

		folderpath := path.Join(root, dirname)
		if _, err := os.Stat(folderpath); os.IsNotExist(err) {
			response.WithError(w, nil, ErrDirectoryNotExist)
			return
		}

		files, err := ioutil.ReadDir(folderpath)
		if err != nil {
			response.WithError(w, nil, ErrDirectoryNotExist)
			return
		}

		data := []*Response{}
		for _, f := range files {
			data = append(data, &Response{
				Filename: f.Name(),
			})
		}

		response.JSON(w, http.StatusOK, data)
	}
}
