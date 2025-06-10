package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/marioanchevski/docu-reach/types"
)

func WriteJsonResponse[T any](w http.ResponseWriter, response types.APIResponse[T]) {
	setJsonHeader(w)
	w.WriteHeader(response.Status)

	if err := json.NewEncoder(w).Encode(response); err != nil {

		log.Printf("Json encoding failed: %v", err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

}

func WriteErrorResponse[T any](w http.ResponseWriter, status int, errorMessage string) {

	response := types.APIResponse[T]{
		Status:    status,
		Timestamp: time.Now(),
		Error:     &errorMessage,
	}

	WriteJsonResponse(w, response)
}

func WriteSuccessResponse[T any](w http.ResponseWriter, status int, message string, data T) {

	response := types.APIResponse[T]{
		Status:    status,
		Timestamp: time.Now(),
		Message:   &message,
		Data:      data,
	}

	WriteJsonResponse(w, response)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func WriteHealthResponse(w http.ResponseWriter) error {
	setJsonHeader(w)
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"status": "UP",
	}
	return json.NewEncoder(w).Encode(response)
}

func setJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
