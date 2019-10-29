package status

import (
	"net/http"

	"github.com/payfazz/chrome-remote-debug/internal/httpserver/response"
)

// GetHandler ...
func GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"status": "ok",
			"code":   http.StatusOK,
		})
	}
}
