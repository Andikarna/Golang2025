package components

import (
	"encoding/json"
	"net/http"
)

// Response format standar
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// RespondJSON untuk response sukses
func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response{
		Status: "success",
		Data:   data,
	}

	json.NewEncoder(w).Encode(resp)
}

// RespondError untuk response error
func RespondError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response{
		Status:  "error",
		Message: message,
	}

	json.NewEncoder(w).Encode(resp)
}
