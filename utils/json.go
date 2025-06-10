package utils

import (
	"encoding/json"
	"net/http"
)

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
