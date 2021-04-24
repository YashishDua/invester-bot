package apis

import (
	"encoding/json"
	"net/http"
)

type HTTPResponse struct {
	StatusCode int         `json:"status_code"`
	Body       interface{} `json:"body"`
}

func HandleResponse(w http.ResponseWriter, payload interface{}) {
	response := HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       payload,
	}

	responseBytes, err := json.Marshal(response)

	if err != nil {
		HandleError(w, InternalServerError(err))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
}
