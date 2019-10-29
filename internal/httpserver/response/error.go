package response

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/payfazz/chrome-remote-debug/pkg/alert"
	"github.com/payfazz/chrome-remote-debug/pkg/errors"
)

// WithError emits proper error response
func WithError(w http.ResponseWriter, n alert.Alert, err error) {
	switch err.(type) {
	case errors.BaseError:
		WithData(w, http.StatusBadRequest, err)
	case errors.NotFoundError:
		WithData(w, http.StatusNotFound, err)
	case errors.CommonError:
		WithData(w, http.StatusBadRequest, err)
	case errors.ValidationError:
		WithData(w, http.StatusUnprocessableEntity, err)
	case errors.AuthError:
		WithData(w, http.StatusUnauthorized, err)
	case errors.PermissionError:
		WithData(w, http.StatusForbidden, err)
	case errors.ServiceError:
		logError(n, err)
		response := errors.NewBaseError(http.StatusText(http.StatusInternalServerError), "Server tidak dapat memproses permintaan anda, cobalah beberapa saat lagi.")
		WithData(w, http.StatusInternalServerError, response)
	default:
		logError(n, err)
		response := errors.NewBaseError(http.StatusText(http.StatusInternalServerError), "Server tidak dapat memproses permintaan anda, cobalah beberapa saat lagi.")
		WithData(w, http.StatusInternalServerError, response)
	}
}

// WithData emits response with data
func WithData(w http.ResponseWriter, status int, data interface{}) {
	JSON(w, status, data)
}

func logError(a alert.Alert, err error) {
	msg := fmt.Sprintf("%+v\n%s", err, string(debug.Stack()))
	log.Println(msg)
	if a != nil {
		message := alert.NewAlertMessage(err.Error(), err, debug.Stack())
		a.Alert(message)
	}
}
