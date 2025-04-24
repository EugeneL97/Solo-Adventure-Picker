package utils

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Error   string `json:"error"`
	Code    int    `json:"code,omitempty"`
	Details string `json:"details,omitempty"`
}

func WriteJSONError(w http.ResponseWriter, status int, error APIError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(error)
}
