package apis

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func BadRequest(err error) HTTPError {
	return HTTPError{
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
	}
}

func InternalServerError(err error) HTTPError {
	return HTTPError{
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	}
}

func HandleError(w http.ResponseWriter, httpError HTTPError) {
	w.WriteHeader(httpError.StatusCode)

	errorBytes, _ := json.Marshal(HTTPError{})

	w.Write(errorBytes)
}
